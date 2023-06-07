#!/bin/bash

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
