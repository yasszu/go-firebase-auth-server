FROM golang:1.16

WORKDIR /go/src/app

RUN go get -u github.com/cosmtrek/air && \
    go get -v github.com/rubenv/sql-migrate/...

EXPOSE 8888

CMD ["air"]