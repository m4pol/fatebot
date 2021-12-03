#!/bin/bash
read -p "Enter your payload name: " pName
go build -ldflags "-s -w" -o bin/$pName cmd/main.go
echo "Build success..."
echo "--------------------------------------------"
echo "Your executable file are in The bin folder."
echo "--------------------------------------------"
