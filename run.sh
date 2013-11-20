#!/usr/bin/env zsh

./clear.sh
go build
Xephyr -ac -screen 800x600 -xinerama :5 &
sleep 1
DISPLAY=:5 ./emptybox
