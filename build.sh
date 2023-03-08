#!/bin/bash
go build -ldflags "-w -s" -o tcpSimulator_amd64 main.go
CGO_ENABLED=0 GOARCH=arm64 GOARM=7 go build -ldflags "-w -s" -o tcpSimulator_arm64 main.go

