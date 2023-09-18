FROM golang:1.17.2-alpine AS build

WORKDIR /src/go-homework
COPY . /src/go-homework
RUN go build -o /bin/go-homework

EXPOSE 9090
ENTRYPOINT ["/bin/go-homework"]