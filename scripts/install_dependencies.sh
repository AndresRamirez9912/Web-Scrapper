#!/bin/bash
apt-get update
go mod tidy
rm app #remove the built app 
chmod +x "$0" # Give permissions to this script 

# Create the env variables from S3
aws s3 cp s3://price-tracker-env/env.json /tmp/env_vars.json

# Read the contents of the JSON file and set environment variables
password=$(jq -r '.EMAIL_PASSWORD' /tmp/env_vars.json) # Replace '.yourKey' with the appropriate JSON key path
my_email=$(jq -r '.MY_EMAIL' /tmp/env_vars.json) # Replace '.yourKey' with the appropriate JSON key path
export MY_EMAIL='$password'
export EMAIL_PASSWORD='$my_email'

# Delete the temporary file
rm /tmp/env_vars.txt
