FROM golang:1.23-alpine

RUN apk add --no-cache git ca-certificates && update-ca-certificates

RUN apk add curl

RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

RUN go get github.com/mattn/go-runewidth

RUN go install github.com/mgechev/revive@latest

RUN go install golang.org/x/tools/cmd/stringer@v0.25.0

