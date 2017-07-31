#!/usr/bin/env bash
 docker run \
    -e "BOT_TOKEN=$1"  \
    -e APP_ENV=testing \
    one-more/acropole