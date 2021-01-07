FROM golang:1.13.4 as builder
ENV DATA_DIRECTORY /go/src/github.com/ejiro-edwin/Financial-App
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -ldflags="-X github.com/ejiro-edwin/Financial-App/internal/config.Version=$APP_VERSION" github.com/ejiro-edwin/Financial-App/cmd/server

FROM alphine:3.10
ENV DATA_DIRECTORY /go/src/github.com/ejiro-edwin/Financial-App
RUN apk add --update --no-cache \
    ca-certificates
COPY --from=builder ${DATA_DIRECTORY}server /Financial-App

ENTRYPOINT ["/Financial-App"]


