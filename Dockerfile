FROM golang:1.26.1-alpine

WORKDIR /app

# Copy the Go module files from the backend project
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the rest of the backend source
COPY backend/ ./

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]