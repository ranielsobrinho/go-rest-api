FROM golang:1.22

#set work directory
WORKDIR /go/src/app

#copy the source code
COPY . .

#expose the port
EXPOSE 5000

#build the Go app
RUN go build -o main cmd/main.go

# run the executable
CMD ["./main"]
