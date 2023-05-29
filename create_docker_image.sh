#!/bin/bash

echo $DOCKER_PASSWORD | docker login --username "$DOCKER_USER" --password-stdin

docker build --no-cache -t alexchub/golang-app:latest .
docker push alexchub/golang-app:latest
docker image rmi alexchub/golang-app:latest
