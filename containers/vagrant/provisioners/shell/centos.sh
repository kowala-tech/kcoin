#!/bin/bash

sudo yum install -y git wget
sudo yum update -y

wget --continue https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz

KUSD_PATH="~vagrant/go/src/github.com/kowala-tech/kcoin/build/bin/"

echo "export PATH=$PATH:/usr/local/go/bin:$KUSD_PATH" >> ~vagrant/.bashrc
