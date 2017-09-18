# Acropole

A Discord bot project to establish a (true) democracy inside a Discord server.

## Getting Started

These instructions will get you a local discord bot running.

### Prerequisites

You just have to have [Docker](https://docs.docker.com/engine/installation/) installed

## Development

To use a development version of Acropole, start by cloning the repository.

```sh
git clone https://gitlab.com/one-more/acropole.git
cd acropole
```

Setup bot token in config.yml file:

```sh
mv app/config.example.yml app/config.yml
nano app/config.yml
```

You have to build development image locally

```sh
docker-compose build
```

You can now run it

```sh
docker-compose up
```

### Testing

You need to have the image build to run tests

```sh
docker build \
    -t one-more/acropole .

 docker run \
    -e APP_ENV=testing \
    one-more/acropole
```

## Work behind proxy

Coming soon