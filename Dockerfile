FROM golang:1.19

WORKDIR /usr/src/gofinance

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/src/gofinance/src/main ./...

CMD ["./main"]
