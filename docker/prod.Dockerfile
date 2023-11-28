FROM golang:alpine AS builder
 
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
 
ADD go.mod .
ADD go.sum .
 
RUN go mod download
RUN go mod tidy

COPY . .
 
RUN go build -ldflags="-s -w" -o main main.go

#======= 

FROM alpine 
 
WORKDIR /app
 
COPY --from=builder /app/main /app/main
COPY --from=builder /app/assets/ ./assets
COPY --from=builder /app/index.html .
 
RUN touch .env // remove it if code wont break
 
EXPOSE 8080
 
CMD ["./main"]