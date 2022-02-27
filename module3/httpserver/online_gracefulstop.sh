#!/bin/sh

PID=`pidof /app/httpserver` && kill -SIGTERM $PID
sleep 5
