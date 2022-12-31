#!/bin/bash

if ! command -v go &> /dev/null || ! command -v upx &> /dev/null ; then
    echo -e "\e[1;31m\"UPX\" packer, or \"Go\" compiler has not been installed!!!\e[0m"
    exit
fi

clear

echo -e "\n\e[1;32mBUILDING x32 DEFAULT...\e[0m\n"
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o bin/$1_x32 cmd/main.go && upx -9 bin/$1_x32 && clear
echo -e "\n\e[1;32mBUILDING x32 MIPS...\e[0m\n"
GOOS=linux GOARCH=mips go build -ldflags "-s -w" -o bin/$1_mips_x32 cmd/main.go && upx -9 bin/$1_mips_x32 && clear
echo -e "\n\e[1;32mBUILDING x32 ARM...\e[0m\n"
GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/$1_arm_x32 cmd/main.go && upx -9 bin/$1_arm_x32 && clear

echo ""
echo -e "\e[1;32m               BUILD SUCCESS!!!                \e[0m"
echo -e "\e[1;37m-----------------------------------------------\e[0m"
echo -e "\e[1;32m  YOUR EXECUTABLE FILES ARE IN THE BIN FOLDER  \e[0m"
echo -e "\e[1;37m-----------------------------------------------\e[0m"
echo ""