#!/bin/bash

docker rm is-tgbot
docker rmi wasd0/is-tgbot

docker build --tag wasd0/is-tgbot -f Dockerfile .
docker push wasd0/is-tgbot:latest 