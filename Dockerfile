FROM golang:alpine as builder

ADD go.mod .

COPY . . 

RUN GOPATH= go build -o /lang-bot ./cmd/lang-bot/main.go

FROM alpine

COPY --from=builder lang-bot lang-bot

CMD ["/lang-bot"]