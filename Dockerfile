FROM golang:latest

ENV APP_HOME /go/src/go-maze

WORKDIR "$APP_HOME"
RUN pwd
RUN echo $(ls -1 )
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY src/go.mod src/go.sum ./
RUN go mod download
RUN go mod verify

COPY src/ .
RUN go build -v -o go-maze

EXPOSE 8000

CMD ["./go-maze"]
