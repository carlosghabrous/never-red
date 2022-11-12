FROM golang:1.17 as go_base

ENV \
    GOPATH=/go \
    APP_ENV=/app \
    PATH=$GOPATH/bin:$APP_ENV:$PATH

FROM go_base as never_red_base

# Create app directory
RUN mkdir /app
WORKDIR /app

ADD src/* .

RUN go mod download && go mod verify

RUN go build -v -o never-red .
EXPOSE 8080
CMD ["/app/never-red"]