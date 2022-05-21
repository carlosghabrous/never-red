FROM golang:1.17 as go_base

ENV \
    GOPATH=/go \
    PATH=$GOPATH/bin:$PATH

FROM go_base as never_red_base
WORKDIR /usr/src/app

COPY go.mod go.su[m] ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin ./...
EXPOSE 9000
CMD ["never-red"]