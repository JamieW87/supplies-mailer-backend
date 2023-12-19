FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/


FROM alpine

WORKDIR /app

RUN apk add --no-cache jq

COPY --from=builder /app/main main
COPY --from=builder /app/supplier-template.html supplier-template.html
COPY ecs-entrypoint.sh .

ENTRYPOINT ./ecs-entrypoint.sh