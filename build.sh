#!/bin/bash

if ! command -v go &> /dev/null || ! command -v upx &> /dev/null ; then
    echo -e "\e[1;31m\"UPX\" packer or \"Go\" compiler is not installed!!!\e[0m"
    exit
fi

go build -ldflags "-s -w" -o bin/$1 cmd/main.go && upx -9 bin/$1 && clear

echo ""
echo -e "\e[1;32m               BUILD SUCCESS!!!               \e[0m"
echo -e "\e[1;37m----------------------------------------------\e[0m"
echo -e "\e[1;32m  YOUR EXECUTABLE FILE ARE IN THE BIN FOLDER  \e[0m"
echo -e "\e[1;37m----------------------------------------------\e[0m"
echo ""
