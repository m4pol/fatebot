#!/bin/bash
go build -ldflags "-s -w" -o bin/$1 cmd/main.go
echo "Build success..."
echo "--------------------------------------------"
echo "Your executable file are in The bin folder."
echo "--------------------------------------------"
