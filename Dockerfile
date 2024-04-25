FROM golang:bookworm

WORKDIR /var/www

COPY .  /var/www

CMD ["go", "run", "main.go"]
