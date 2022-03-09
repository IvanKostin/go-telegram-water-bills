FROM golang:1.16-alpine3.15
ENV TELEGRAM_APITOKEN=""

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD [ "/main" ]
