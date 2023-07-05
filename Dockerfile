# syntax = docker/dockerfile:1-experimental

# Build the manager binary
FROM golang:1.19.10-alpine3.18 as builder
ARG CRYPTO_LIB
ENV GOEXPERIMENT=${CRYPTO_LIB:+boringcrypto}

RUN apk update
RUN apk add git gcc g++ curl

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/

# Build

RUN --mount=type=cache,target=/root/.cache/go-build \
    if [ ${CRYPTO_LIB} ]; \
    then \
      CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags "-linkmode=external -extldflags=-static" -o manager main.go ;\
    else \
      CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o manager main.go ;\
    fi
       
# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
