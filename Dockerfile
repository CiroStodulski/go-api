FROM golang:1.16-alpine

WORKDIR /app
EXPOSE 8080

COPY dist .
#COPY go-api .

ENTRYPOINT ["./go-api"]