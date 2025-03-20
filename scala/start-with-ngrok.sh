#!/bin/bash

# Start Play Framework app in the background
nohup sbt run &

# Start ngrok
ngrok http 9000

