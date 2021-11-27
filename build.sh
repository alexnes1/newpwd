#!/usr/bin/env bash

VERSION=$1
BIN_PATH=./bin/
for OS in linux windows
do 
    for ARCH in 386 amd64 arm arm64
    do    
        EXE_PATH="$BIN_PATH/$OS.$ARCH"
        mkdir -p $EXE_PATH
        echo "Building $OS/$ARCH..."        
        if [ $OS == "windows" ]; then 
            FILENAME="newpwd.exe"
        else
            FILENAME="newpwd"
        fi
        CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w -X 'main.version=$VERSION'" -o $EXE_PATH/$FILENAME ./cmd/newpwd/main.go  && upx --best $EXE_PATH/$FILENAME
        tar -czf "${BIN_PATH}/newpwd_${OS}_${ARCH}_${VERSION}.tgz" -C $EXE_PATH $FILENAME
    done
done