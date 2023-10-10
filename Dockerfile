FROM golang:1.21

WORKDIR /app

ARG DEFAULT_URLS="https://toscrape.com/,https://www.scrapethissite.com/pages/,https://example.com/"

ENV URLS=${DEFAULT_URLS}

COPY . .

RUN go mod download

RUN go build -o webscraper
# Define an environment variable for passing URLs to the container
#ENV URLS https://toscrape.com/,https://www.scrapethissite.com/pages/,https://example.com/

CMD ["./webscraper"]