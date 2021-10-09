FROM golang:1.16-alpine

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["/main"]