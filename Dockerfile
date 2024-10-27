FROM golang:1.23.1 AS builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY /. .
COPY .env .

RUN go build -o main.out cmd/is-tgbot/main.go

FROM alpine

COPY --from=builder /app/main.out /main.out
COPY --from=builder /app/.env /.env
COPY --from=builder /app/config /config

CMD ["apk add --no-cache ca-certificates"]
ENTRYPOINT ["/main.out"]