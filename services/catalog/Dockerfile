FROM golang:1.17-alpine AS builder

RUN apk add --no-cache git

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN  CGO_ENABLED=0 GOOS=linux go build -o usr/bin/catalog ./services/catalog
FROM alpine

EXPOSE 8080
WORKDIR /root/

COPY --from=builder work/usr/bin/catalog /usr/bin/catalog
CMD ["/usr/bin/catalog"]
