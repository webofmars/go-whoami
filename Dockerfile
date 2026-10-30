FROM golang:alpine

WORKDIR /go/src/app
ADD . /go/src/app/

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

EXPOSE 3000