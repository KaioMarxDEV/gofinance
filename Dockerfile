FROM golang:1.19

WORKDIR /gofinance

COPY go.mod /gofinance/
COPY go.mod /gofinance/
RUN go mod download

COPY . /gofinance/
RUN go build -o /server

EXPOSE 3000

CMD [ "/server" ]
