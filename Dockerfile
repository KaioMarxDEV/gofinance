FROM golang:1.19 AS backend-builder

LABEL app="Gofinance"
LABEL author="Kaio Marx. <kaiomarxdev@gmail.com>"
LABEL description="Finance managment application"

WORKDIR /backendcompile

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./server ./main.go

EXPOSE 3000
CMD [ "./server" ]
