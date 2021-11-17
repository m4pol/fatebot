#!/bin/bash
read -p "Enter your payload name: " pName
go build -ldflags "-s -w" -o $pName main.go
echo "Build success..."
echo "This payload is for upload on FTP server."
