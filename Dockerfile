FROM golang:1.21

WORKDIR /app

ENV URLS=""

COPY . .

RUN go mod download

RUN go build -o webscraper

CMD ["./webscraper"]