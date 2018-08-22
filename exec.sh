#!/usr/bin/env bash
docker build -t upworkdemo .
docker run --env-file .env -p 3000:3000 -it upworkdemo
