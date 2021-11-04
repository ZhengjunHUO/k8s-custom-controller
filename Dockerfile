FROM golang:1.16.7-alpine3.13 AS build
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY pkg/ ./pkg/
COPY main.go .
RUN CGO_ENABLED=0 go build -o custom-controller .

FROM alpine:3.13.6
LABEL maintainer="Zhengjun HUO"
COPY --from=build /workspace/custom-controller /usr/local/bin/custom-controller
ENTRYPOINT ["/usr/local/bin/custom-controller"]
