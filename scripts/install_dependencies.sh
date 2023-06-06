#!/bin/bash

# Install dependencies
sudo apt-get update
sudo apt install jq

# Go to the application path
cd home/ubuntu/price-tracking/Web-Scrapper/

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
