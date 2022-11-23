FROM golang:1.19

WORKDIR /go/src/github.com/KaioMarxDEV/gofinance

COPY go.mod /go/src/github.com/KaioMarxDEV/gofinance/
COPY go.mod /go/src/github.com/KaioMarxDEV/gofinance/
RUN go mod download

COPY . /go/src/github.com/KaioMarxDEV/gofinance/
RUN go build -o /server

EXPOSE 3000

CMD [ "/server" ]
