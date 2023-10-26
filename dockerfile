FROM golang:1.20.10-alpine3.17

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY main.go .
COPY main_test.go .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]