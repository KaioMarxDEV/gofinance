FROM golang:1.19 AS backend-builder

LABEL app="Gofinance"
LABEL author="Kaio Marx. <kaiomarxdev@gmail.com>"
LABEL description="Finance managment application"

WORKDIR /backendcompile

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./server ./main.go

FROM node:18 AS frontend-builder

WORKDIR /frontendcompile

COPY ./web/package.json ./web/yarn.lock ./

RUN yarn

COPY ./web .

RUN yarn build

FROM alpine:latest AS prod

WORKDIR /build

RUN mkdir -p /web/build

COPY --from=backend-builder /backendcompile/server .
COPY --from=frontend-builder /frontendcompile/dist ./web/build

EXPOSE 3000
ENTRYPOINT [ "./server" ]
