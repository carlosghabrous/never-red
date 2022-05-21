FROM golang:1.17

WORKDIR /usr/src/app
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

COPY go.mod go.su[m] ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin ./...

EXPOSE 9000

CMD ["never-red"]