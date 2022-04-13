FROM golang:1.18.0-alpine3.15

WORKDIR /app

ADD . .

RUN go build main.go

CMD ["/app/main"]
