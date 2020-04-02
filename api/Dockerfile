FROM golang:latest AS builder

RUN mkdir /go/src/go_mongodb/
WORKDIR /go/src/go_mongodb/api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/go_mongodb/api .

EXPOSE 8080
CMD ["./api"]