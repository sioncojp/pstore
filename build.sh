#!/bin/bash -
declare -r Name="pstore"

for GOOS in darwin linux; do
    GO111MODULE=on GOOS=$GOOS GOARCH=amd64 go build -o bin/pstore-$GOOS-amd64 *.go
done
