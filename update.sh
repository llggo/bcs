#!/bin/sh

pm2 stop qrcode-pba
rm qrcode-pba.exe
go build -o qrcode-pba.exe
pm2 start qrcode-pba.exe

while true; do
    
    if [[ $(git pull origin master) == *up-to-date* ]]; 
    then
        echo "no change"
    else
        echo "detect changes"
        sleep 2s
        echo "stop app"
		pm2 stop qrcode-pba
		rm qrcode-pba.exe
        go build -o qrcode-pba.exe
        pm2 restart qrcode-pba
    fi

    echo "sleep 30s"

    sleep 30s        

done
