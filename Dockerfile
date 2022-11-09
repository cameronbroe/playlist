FROM golang:1.18.4 AS build

ENV CGO_ENABLED=0

WORKDIR /playlist-api
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY . .
RUN go build -o build/server cmd/server.go

FROM debian:11-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=build /playlist-api/build/server /server

EXPOSE 8080

CMD ["/server"]
