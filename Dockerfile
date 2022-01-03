FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o voiddb voiddb.go
EXPOSE 8080
CMD ["./voiddb"]