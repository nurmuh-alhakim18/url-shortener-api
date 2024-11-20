FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY cmd/url-shortener /usr/bin/url-shortener

CMD [ "url-shortener" ]