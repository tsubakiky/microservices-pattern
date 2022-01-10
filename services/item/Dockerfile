FROM golang:1.17-alpine AS builder

RUN apk add --no-cache git

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN  CGO_ENABLED=0 GOOS=linux go build -o usr/bin/item ./services/item
FROM alpine

EXPOSE 8080
WORKDIR /root/

COPY --from=builder work/usr/bin/item /usr/bin/item
CMD ["/usr/bin/item"]
