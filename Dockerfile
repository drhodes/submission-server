FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD *.go /app/
WORKDIR /app
RUN apk add git # try to not need the next three lines
RUN go get "github.com/abbot/go-http-auth"
RUN go get "golang.org/x/crypto/bcrypt"

RUN go build -o submitter . 
ENTRYPOINT ["/app/submitter"]
