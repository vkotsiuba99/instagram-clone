FROM golang:1.18-alpine3.16

COPY ./internal /go/src/instagram-clone/account-api

WORKDIR /go/src/instagram-clone/account-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD [ "account-api" ]