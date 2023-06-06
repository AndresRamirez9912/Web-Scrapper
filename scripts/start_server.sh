#!/bin/bash

# Go to the application path
cd price-tracking/Web-Scrapper/

# Build the applicatiojn
go build -o ./app

# Start the server
./app
