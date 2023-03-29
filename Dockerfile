# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN curl -L https://github.com/swaggo/swag/releases/download/v1.8.12/swag_1.8.12_Linux_x86_64.tar.gz | tar xvz
RUN ./swag init
RUN go build -o main main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main . 
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./db/migrations

EXPOSE 8080
CMD [ "/app/main" ]
RUN chmod +x /app/start.sh
RUN chmod +x /app/wait-for.sh
ENTRYPOINT [ "/app/start.sh" ]