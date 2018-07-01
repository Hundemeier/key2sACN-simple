#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m' # No Color
#TICK='\033[0;32m \xE2\x9C\x93 \033[0m'
TICK=$GREEN'\xE2\x9C\x93'$NC

printf 'building for...\n'
printf 'Linux AMD64'
if GOOS=linux GOARCH=amd64 go build -o build/key2sACN-linux-64; then 
  printf ' '$TICK'\n' $TICK
fi

printf 'Linux 386'
if GOOS=linux GOARCH=386 go build -o build/key2sACN-linux-32; then
  printf ' '$TICK'\n' $TICK
fi

printf 'Linux ARM'
if GOOS=linux GOARCH=arm go build -o build/key2sACN-linux-arm; then
  printf ' '$TICK'\n' $TICK
fi

printf 'Windows AMD64'
if GOOS=windows GOARCH=amd64 go build -o build/key2sACN-windows-64.exe; then
  printf ' '$TICK'\n' $TICK
fi

printf 'Windows 386'
if GOOS=windows GOARCH=386 go build -o build/key2sACN-windows-32.exe; then
  printf ' '$TICK'\n' $TICK
fi

#printf 'Darwin AMD64'
#GOOS=darwin GOARCH=amd64 go build -o build/key2sACN-macos-64

#printf 'Darwin 386'
#GOOS=darwin GOARCH=386 go build -o build/key2sACN-macos-32