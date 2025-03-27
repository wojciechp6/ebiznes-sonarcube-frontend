#!/bin/bash

docker run -d -p 9000:9000 -t wojciechp6/scala 
ngrok http 9000

