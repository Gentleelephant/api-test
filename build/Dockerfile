FROM golang:latest

WORKDIR /app

ADD . /app

RUN go build -o api-test ./main.go

EXPOSE 9090

CMD ["sh"]