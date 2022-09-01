FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o tictactoe cmd/main.go

FROM alpine

COPY --from=builder /build/tictactoe .

CMD "./tictactoe"