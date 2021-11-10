#!/bin/bash

cd cmd
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -ldflags "-s -w" -o ../dist/nectar_darwin_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -ldflags "-s -w" -o ../dist/nectar_window_386.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -ldflags "-s -w" -o ../dist/nectar_window_amd64.exe
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build  -ldflags "-s -w" -o ../dist/nectar_linux_386
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w" -o ../dist/nectar_linux_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ../dist/nectar_linux_arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ../dist/nectar_linux_arm64
