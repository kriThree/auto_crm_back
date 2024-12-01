FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/crm
EXPOSE 5001
CMD ["./main "]

