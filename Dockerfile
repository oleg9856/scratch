FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git curl build-base
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
WORKDIR /app
COPY . .
COPY run_goose.sh ./run_goose.sh
RUN chmod +x ./run_goose.sh
RUN sqlc generate
RUN go build -o rssagg .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rssagg .
COPY --from=builder /app/run_goose.sh ./run_goose.sh
EXPOSE 8082
ENV PORT=8082
CMD ["./rssagg"]
