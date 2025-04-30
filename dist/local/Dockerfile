FROM alpine:latest
WORKDIR /app

COPY bin/bootstrapper .

EXPOSE 8080
ENTRYPOINT ["/app/bootstrapper"]