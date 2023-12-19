#Build stage
FROM golang:1.21.5 as builder
WORKDIR /app
COPY . .
# Effectively tracks changes within your go.mod file
COPY go.mod .
 
RUN go mod download

RUN go build -o main main.go

#Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]