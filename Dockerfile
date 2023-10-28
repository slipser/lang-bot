FROM golang:alpine as builder

WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN GOPATH= go build -o /lang-bot cmd/lang-bot/main.go

FROM alpine

COPY --from=builder lang-bot /build/lang-bot

CMD ["/build/lang-bot"]