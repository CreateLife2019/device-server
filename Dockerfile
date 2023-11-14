# syntax=docker/dockerfile:1

FROM --platform=$TARGETPLATFORM golang:alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

COPY . ./
COPY *.json ./
COPY *.html ./
RUN export GOPROXY=https://goproxy.cn
RUN go mod tidy
RUN go build -o /device-server

FROM alpine AS runner
WORKDIR /app
COPY --from=builder . ./
COPY --from=builder *.json ./
COPY --from=builder *.html ./

EXPOSE 8090

CMD [ "/device-server" ]
