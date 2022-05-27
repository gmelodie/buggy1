#!/bin/bash

# init script for interview box
apt update
apt install -y golang-go git plocate

updatedb

git clone https://github.com/gmelodie/buggy1 /opt/api-server
cd /opt/api-server
go build && mv buggy1 api-server && ./api-server &

