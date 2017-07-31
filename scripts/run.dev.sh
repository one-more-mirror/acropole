#!/usr/bin/env bash

docker kill acropole
docker rm acropole

docker run \
    -it \
    -e "BOT_TOKEN=$1"  \
    --name acropole \
    one-more/acropole