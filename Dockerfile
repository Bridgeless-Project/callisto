FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/forbole/callisto

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"
ENV GOPRIVATE=github.com/*
ENV GONOSUMDB=github.com/*
ENV GONOPROXY=github.com/*

COPY ./go.mod ./go.sum ./
# Read the CI_ACCESS_TOKEN from the .env file
ARG CI_ACCESS_TOKEN
RUN git config --global url."https://olegfomenkodev:${CI_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"
COPY . .

RUN #go mod vendor
RUN go build -mod=vendor -o /usr/local/bin/bdjuno /go/src/github.com/forbole/callisto/cmd/bdjuno


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/bdjuno /usr/local/bin/bdjuno

RUN apk add --no-cache ca-certificates

COPY ./genesis.json /genesis/genesis.json

ENTRYPOINT ["bdjuno"]