#!/bin/bash

# Instal Go
# Install dependencies
rm go1.20.5.linux-amd64.tar.gz
curl -O https://storage.googleapis.com/golang/go1.20.5.linux-amd64.tar.gz #Download the latest version 
tar -xvf go1.20.5.linux-amd64.tar.gz # Extract the tar
sudo mv go /usr/local
echo "export GOPATH=$HOME/work" >> ~/.profile 
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.profile
source ~/.profile
go version

# Go to the application path
cd /home/ubuntu/price-tracking/Web-Scrapper/

# Remove the previously files
rm app

# Build the applicatiojn
go build -o /home/ubuntu/price-tracking/Web-Scrapper/app > /home/ubuntu/price-tracking/priceTracking.log 
sleep 10 # Delay meanwhile the built is created
if [ -f /home/ubuntu/price-tracking/Web-Scrapper/app ]; then
    echo "File exists"
else
    echo "File does not exist"
    exit 1
fi

# Start the server in the background
/home/ubuntu/price-tracking/Web-Scrapper/app &
