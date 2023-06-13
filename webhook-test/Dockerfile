FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENTRYPOINT ["/app/server"]
