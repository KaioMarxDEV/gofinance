FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN go build ./src/main/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/main"]

