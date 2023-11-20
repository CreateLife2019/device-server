# syntax=docker/dockerfile:1
ARG  TARGETPLATFORM=linux/amd64 
FROM --platform=${TARGETPLATFORM} golang:alpine AS builder
USER root
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

COPY . ./
RUN export GOPROXY=https://goproxy.io,direct
RUN go mod tidy
RUN go build -o /device-server

FROM alpine AS runner
WORKDIR /app
COPY --from=builder . ./

EXPOSE 8090

CMD [ "/device-server" ]
