#!/bin/bash
apt-get update
go mod tidy
rm app #remove the built app 
chmod +x "$0" # Give permissions to this script 
