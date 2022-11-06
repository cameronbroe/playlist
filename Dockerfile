FROM golang:1.18.4 AS build

ENV CGO_ENABLED=0

WORKDIR /music.cameronbroe.com
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY . .
RUN go build -o build/server cmd/server.go

FROM scratch
COPY --from=build /music.cameronbroe.com/build/server /server

EXPOSE 8080

CMD ["/server"]
