#Build stage
FROM golang:alpine as builder
WORKDIR /app
COPY . .

RUN go build -o main main.go

#Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]