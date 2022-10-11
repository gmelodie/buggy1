#!/bin/bash

# init script for interview box
apt update
apt install -y golang-go git plocate make

updatedb

go install github.com/swaggo/swag/cmd/swag@latest
echo 'export PATH=$PATH:$HOME/go/bin' >> /root/.bashrc # not sure if this persists

git clone https://github.com/gmelodie/buggy1 /opt/api-server
cd /opt/api-server
go build && mv buggy1 api-server && ./api-server &
