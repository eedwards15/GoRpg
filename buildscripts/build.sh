#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o bin/windows/RPG.exe ../main.go
GOOS=darwin GOARCH=arm64 go build -o bin/mac/RPG ../main.go