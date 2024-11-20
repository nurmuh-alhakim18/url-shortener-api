FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C cmd -o url-shortener

ENV HOST=0.0.0.0

CMD [ "/app/cmd/url-shortener" ]