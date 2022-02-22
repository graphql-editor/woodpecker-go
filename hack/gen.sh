#!/bin/sh

docker run --rm -it  --user "$(id -u)":"$(id -g)" -e GOPATH="$(go env GOPATH)":/go -v "$HOME":"$HOME" -w "$(pwd)" quay.io/goswagger/swagger generate client -f ./swagger.yaml
