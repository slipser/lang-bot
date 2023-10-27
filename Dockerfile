FROM golang:alpine

WORKDIR /build

COPY . . 

RUN go build -o /lang-bot ./cmd/lang-bot/main.go

CMD ["/lang-bot"]