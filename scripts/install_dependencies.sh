#!/bin/bash

# Install dependencies
rm go1.20.5.linux-amd64.tar.gz
curl -O https://storage.googleapis.com/golang/go1.20.5.linux-amd64.tar.gz #Download the latest version 
tar -xvf go1.20.5.linux-amd64.tar.gz # Extract the tar
sudo mv go /usr/local
echo "export GOPATH=$HOME/work" >> ~/.profile 
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.profile
source ~/.profile
go version

sudo apt-get update
sudo apt install jq

# Go to the application path
cd /home/ubuntu/price-tracking/Web-Scrapper/

# Install Golang pacakages 
go mod tidy

# Create the env variables from S3
aws s3 cp s3://price-tracker-env/env.json /tmp/env_vars.json

# Read the contents of the JSON file and set environment variables
EMAIL_PASSWORD=$(jq -r '.EMAIL_PASSWORD' /tmp/env_vars.json)
MY_EMAIL=$(jq -r '.MY_EMAIL' /tmp/env_vars.json) 

#Create the env variables
export EMAIL_PASSWORD 
export MY_EMAIL

# Create .env file 
sudo touch .env

# Write the env variables on the .env file
sudo sh -c 'echo "EMAIL_PASSWORD=$EMAIL_PASSWORD" >> .env'
sudo sh -c 'echo "MY_EMAIL=$MY_EMAIL" >> .env'

# Delete the temporary file
rm /tmp/env_vars.json

# Build application
cd /home/ubuntu/price-tracking/Web-Scrapper/

# Remove the previously files
rm app
rm nohup.out

# Build the applicatiojn
go build -o /home/ubuntu/price-tracking/Web-Scrapper/app > /home/ubuntu/price-tracking/priceTracking.log 
