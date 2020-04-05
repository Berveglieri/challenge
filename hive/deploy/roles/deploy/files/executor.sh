#!/usr/bin/env bash

IMAGENAME="017136389210.dkr.ecr.eu-central-1.amazonaws.com/web_app:v1.0"
NAME=$(docker ps | awk '{print $14}')

if [ -z "$NAME" ]
then
    eval "$(aws ecr get-login --no-include-email --region eu-central-1)"
    docker pull $IMAGENAME
    docker run -d -p 3333:3333 $IMAGENAME
else
    docker rm $NAME --force
    eval "$(aws ecr get-login --no-include-email --region eu-central-1)"
    docker pull $IMAGENAME
    docker run -d -p 3333:3333 $IMAGENAME
fi

