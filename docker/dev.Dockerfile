FROM golang:latest
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT ["air"]