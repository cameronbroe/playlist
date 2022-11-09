# playlist.cameronbroe.com

This repository contains the Go API and React application for consuming my mega-playlist. The motivation behind this application is to act as a place for me to dump tracks I like and easily get cross-service links to share with friends who use different streaming services than I do. It uses iTunes Search to decorate an object consisting of the song title, artist, and album with the Apple Music URL. It then uses Odesli's API to translate those Apple Music URLs into Spotify and YouTube URLs. I have Shortcuts on my devices that allow me to easily submit whatever I am currently listening to the API.

# Running the API

The API is built using Go with Go Modules and uses a pure Go port of SQLite for its backing database. This allows the application to be built into purely static binaries that can easily just run wherever a Go toolchain can be installed. I am using a Raspberry Pi 4 to work on this project. There is also a Dockerfile for those who prefer to use Docker.

## Go toolchain

Build the API binary:

`go build -o build/playlist-api cmd/server.go`

And run the API:

`build/playlist-api`

The database will be created in the current working directory with the name `database.sqlite`.

There is also support for requiring clients to provide a static API key when submitting. Just pass in the static key through the environment:

`API_SECRET=hunter2 build/playlist-api`

You can also configure a different location for the database file on the filesystem:

`DATABASE_PATH=/path/to/database.sqlite build/playlist-api`

## Docker

Build the image using the Dockerfile:

`docker build -t playlist-api .`

And run it with the port forwarded:

`docker run --rm -p 8080:8080 playlist-api`

Or if you want data to persist, then use a volume mount and specify a custom database path:

`docker run --rm -p 8080:8080 -v $(pwd)/dbdata:/dbdata -e DATABASE_PATH=/dbdata/database.sqlite playlist-api`

## Front-end

The front-end is a basic React app that was bootstrapped with create-react-app. Follow the README.md in the `frontend/` folder.

## License

This repo is MIT licensed, check the `LICENSE` file for more information.