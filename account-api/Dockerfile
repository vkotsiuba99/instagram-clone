FROM golang:1.18-alpine3.16 as builder

ADD . /go/src/instagram-clone/

WORKDIR /go/src/instagram-clone/account-api

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /account-api . -ldflags "-X main.env=docker"

FROM alpine

RUN apk update \
  && apk add --no-cache

COPY --from=builder /account-api ./
RUN chmod +x ./account-api
CMD [ "./account-api" ]