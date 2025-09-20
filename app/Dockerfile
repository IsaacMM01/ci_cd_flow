FROM golang:alpine

WORKDIR /APP_1
COPY go_app.go .

RUN go build -o go_app go_app.go

EXPOSE 8080

CMD ["./go_app"]


