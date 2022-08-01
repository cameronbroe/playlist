FROM golang:1.18.4 AS build
WORKDIR /music.cameronbroe.com
COPY . .
RUN go mod download
RUN go build -o build/server cmd/server.go

FROM scratch
COPY --from=build /music.cameronbroe.com/build/server /server
CMD ["/server"]