FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build src/main/main.go

EXPOSE 3000

CMD ["./main"]
