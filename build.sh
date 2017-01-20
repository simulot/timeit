#!/bin/sh

for i in "linux" "windows"
do 
    mkdir -p binaries/$i
    exe=timeit
    if [ $i = "windows" ]  
    then
        exe=timeit.exe
    fi
    GOOS=$i go build -o binaries/$i/$exe *.go
done

