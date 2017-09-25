#!/bin/sh
rm qrcode-pba.exe
go build -o qrcode-pba.exe
qrcode -conf=qrcode-pba.toml