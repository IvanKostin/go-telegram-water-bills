FROM golang:1.16-alpine3.15
ARG token
ENV TELEGRAM_APITOKEN=$token

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD [ "/main" ]
