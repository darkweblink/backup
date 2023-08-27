#!/bin/bash
# 综合性脚本，综合build、build-all、build-linux、build-win等操作

# 编译到 win 和 linux 平台
function build(){
    build_linux;
    build_win;
    echo "编译完成"
}

# 编译到 linux 平台
function build_linux(){
    echo "开始编译到 Linux X64平台"
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go_gatway main.go
}

# 编译到 win 平台
function build_win(){
    echo "开始编译到 Windows X64 平台"
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o go_gatway.exe main.go
}

# 编译到全平台
function build_all(){
    build_linux;
    mv go_gatway go_gatway-linux-x86_64
    echo "开始编译到 Linux X86 平台"
    CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o go_gatway-linux-x86 main.go
    echo "开始编译到 Linux ARM v5 平台"
    CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o go_gatway-linux-armv5 main.go
    echo "开始编译到 Linux ARM v6 平台"
    CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-s -w" -o go_gatway-linux-armv6 main.go
    echo "开始编译到 Linux ARM v7 平台"
    CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o go_gatway-linux-armv7 main.go
    build_win;
    mv go_gatway.exe go_gatway-win-x86_64.exe
    echo "编译完成"
}

# 帮助文档
function help(){
    echo ""
    echo "Usage: ./bash.sh [OPTION=help]"
    echo ""
    echo "OPTIONS: "
    echo "  build       编译到 windows 和 Linux 平台【 x86_64 架构】"
    echo "  build-all   编译到 windows 和 Linux 平台的全架构"
    echo "  build-linux 编译到 Linux 系统的 x86_64 架构"
    echo "  build-win   编译到 Windows 系统的 x86_64 架构"
    echo "  help        输出本帮助信息"
}

# 根据用户输入的option进行选择执行的命令
case $1 in 
build)
    build;;
build-linux)
    build_linux;
    echo "编译完成";;
build-win)
    build_win;
    echo "编译完成";;
build-all)
    build_all;;
*)
    help;;
esac
