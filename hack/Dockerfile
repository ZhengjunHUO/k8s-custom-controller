FROM golang:1.18.4

ENV GO111MODULE=off
RUN go get k8s.io/code-generator; exit 0
RUN go get k8s.io/apimachinery; exit 0
ARG REPO_NAME="github.com/<PATH_TO_REPO>"
ARG PATH_TO_REPO="${GOPATH}/src/${REPO_NAME}"
RUN mkdir -p $PATH_TO_REPO
VOLUME $PATH_TO_REPO
WORKDIR $GOPATH/src/k8s.io/code-generator
