FROM golang:latest

ENV APP_HOME /go/src/client-server

WORKDIR "$APP_HOME"
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go mod verify

COPY ./ .
RUN go build -v -o client-server

EXPOSE 7000

CMD ["./client-server"]
