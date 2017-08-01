# Acropole

A Discord bot project to establish a (true) democracy inside a Discord server.

## Getting Started

These instructions will get you a local discord bot running.

### Prerequisites

You just have to have [Docker](https://docs.docker.com/engine/installation/) installed

### Usage

```sh
docker run registry.gitlab.com/one-more/acropole:master -e BOT_TOKEN="Bot xxxxxxx"
```

## Development

To use a development version of Acropole, start by cloning the repository.

```sh
git clone https://gitlab.com/one-more/acropole.git
cd acropole
```

You have to build development image locally

```sh
docker build \
    --build-arg app_env=development \
    -t one-more/acropole .
```

You can now run it

```sh
docker run \
    -it \
    -e "BOT_TOKEN=Bot xxxxxxxxxxxxxxxxxxxxxxx"  \
    --name acropole \
    one-more/acropole
```

### Testing

You need to have the image build to run tests

```sh
 docker run \
    -e "BOT_TOKEN=Bot xxxxxxxxxxxxxxxxxxxxxxx"  \
    -e APP_ENV=testing \
    one-more/acropole
```

## Work behind proxy

### Build

To build it behind with proxy:

Build command: 
```sh
docker build \
    --build-arg "HTTP_PROXY=<PROXY_URL>" \
    --build-arg "HTTPS_PROXY=<PROXY_URL>" \
    --build-arg "http_proxy=<PROXY_URL>" \
    --build-arg "https_proxy=<PROXY_URL>" \
    --build-arg app_env=development \
    -t one-more/acropole .
```

Run command:
```sh
docker run \
    -it \
    -e "HTTP_PROXY=<PROXY_URL>" \
    -e "HTTPS_PROXY=<PROXY_URL>" \
    -e "http_proxy=<PROXY_URL>" \
    -e "https_proxy=<PROXY_URL>" \
    -e "BOT_TOKEN=Bot xxxxxxxxxxxxxxxxxxxxxxx"  \
    --name acropole \
    one-more/acropole
```

### Test

To run tests it behind with proxy:

```sh
 docker run \
    -e "BOT_TOKEN=Bot xxxxxxxxxxxxxxxxxxxxxxx"  \
    -e "HTTP_PROXY=<PROXY_URL>" \
    -e "HTTPS_PROXY=<PROXY_URL>" \
    -e "http_proxy=<PROXY_URL>" \
    -e "https_proxy=<PROXY_URL>" \
    -e APP_ENV=testing \
    one-more/acropole
```