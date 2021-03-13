FROM golang:latest

LABEL maintainer="GUEKENG TINDO Yvan" version="1.0"

WORKDIR $GOPATH/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app

EXPOSE 8000

CMD ["./app"]