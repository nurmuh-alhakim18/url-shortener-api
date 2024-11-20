FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd -o url-shortener

EXPOSE 8080

CMD [ "./cmd/url-shortener" ]