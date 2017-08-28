#!/bin/sh
rm bqrc.exe
go build -o bqrc.exe
bqrc -conf=bqrc.toml