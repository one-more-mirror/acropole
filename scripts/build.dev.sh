#!/usr/bin/env bash
docker build \
    --build-arg app_env=development \
    -t one-more/acropole .
