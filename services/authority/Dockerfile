FROM golang:1.17-alpine AS builder

RUN apk add --no-cache git

WORKDIR /work

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/authority ./services/authority
FROM alpine

EXPOSE 8080
WORKDIR /root/

COPY --from=builder /usr/bin/authority .
CMD ["/usr/bin/authority"]
