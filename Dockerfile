FROM golang:1.18

COPY . /app
WORKDIR /app

EXPOSE 8080

#RUN CGO_ENABLED=0 GOOS=linux go build -o main
RUN CGO_ENABLED=0 GOOS=linux go build -o main

CMD ["./main"]
