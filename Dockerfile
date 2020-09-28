FROM golang:latest

RUN mkdir /app
COPY ./src /app

WORKDIR /app

EXPOSE 80

ENTRYPOINT [ "go", "run", "." ]