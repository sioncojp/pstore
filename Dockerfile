# build stage
FROM golang:1.12.10 AS builder

ENV GO111MODULE auto
ENV CGO_ENABLED=0

ADD . /src
WORKDIR /src
RUN make build

# final stage
FROM alpine

ENV AWS_SDK_LOAD_CONFIG=1

WORKDIR /app
COPY --from=builder /src/bin/pstore /app/
