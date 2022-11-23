FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN go build ./src/main/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /root

COPY --from=build /app ./

EXPOSE 3000

ENTRYPOINT ["./main"]

