#!/bin/bash

export GOLANG_VERSION="1.16.5"
# install Go
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go{$GOLANG_VERSION}.linux-amd64.tar.gz
#export go to environment varible
export PATH=$PATH:/usr/local/go/bin
#check version
go version
#Install gcc compiler Ubuntu
sudo apt-get install build-essential
