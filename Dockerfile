FROM golang:1.20.4 AS builder

WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS runner

COPY --from=builder /usr/src/app .

CMD ["./app"]