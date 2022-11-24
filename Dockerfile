FROM golang:1.17 as go_base

# Create app directory
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY src/ ./src
RUN go build -v -o . ./...

EXPOSE 8080
CMD ["./never-red"]
