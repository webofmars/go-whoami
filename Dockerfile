FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app .

FROM scratch

WORKDIR /app
COPY --from=builder /app/app ./
COPY templates ./templates
COPY html ./html

CMD ["./app"]

EXPOSE 3000
