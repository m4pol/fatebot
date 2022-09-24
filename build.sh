#!/bin/bash

if ! command -v go &> /dev/null || ! command -v upx &> /dev/null ; then
    echo -e "\e[1;31m\"UPX\" packer or \"Go\" compiler has not installed!!!\e[0m"
    exit
fi

clear

GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o bin/$1_x32 cmd/main.go && upx -9 bin/$1_x32 && clear
GOOS=linux GOARCH=mips go build -ldflags "-s -w" -o bin/$1_mips_x32 cmd/main.go && upx -9 bin/$1_mips_x32 && clear

# If you want to use the pure x64 executable just uncomment it, but idk why upx can't compress the mips x64 executable file.

# GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/$1_x64 cmd/main.go && upx -9 bin/$1_x64 && clear
# GOOS=linux GOARCH=mips64 go build -ldflags "-s -w" -o bin/$1_mips_x64 cmd/main.go && upx -9 bin/$1_mips_x64 && clear

echo ""
echo -e "\e[1;32m               BUILD SUCCESS!!!                \e[0m"
echo -e "\e[1;37m-----------------------------------------------\e[0m"
echo -e "\e[1;32m  YOUR EXECUTABLE FILEs ARE IN THE BIN FOLDER  \e[0m"
echo -e "\e[1;37m-----------------------------------------------\e[0m"
echo ""
