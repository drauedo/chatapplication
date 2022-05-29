FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod init chat
RUN go get -t
RUN go build -o main .

EXPOSE 8888

CMD ["/app/main"]