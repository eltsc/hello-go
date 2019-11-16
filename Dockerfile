FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/go-sql-driver/mysql
RUN go build serverSQL.go

CMD [ "./serverSQL" ]

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["app"]

EXPOSE 8080