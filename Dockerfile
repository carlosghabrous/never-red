FROM golang:1.18 as go_base

# Create app directory
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY .env .env
COPY src/ ./src
RUN go build -v -o . ./...

EXPOSE 8082
CMD ["./never-red"]
