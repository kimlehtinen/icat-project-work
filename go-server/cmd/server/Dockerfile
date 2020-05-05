FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

# on save: auto-compile and run server
ENTRYPOINT CompileDaemon --build="go build -o icat ./cmd/server/" --command="./icat"