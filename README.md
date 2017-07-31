# Acropole

A Discord bot project to establish a (true) democracy inside a Discord server.

## Getting Started

These instructions will get you a local discord bot running.

### Prerequisites

You just have to have [Docker](https://docs.docker.com/engine/installation/) installed

### Installation

```
docker run registry.gitlab.com/one-more/acropole:master -e BOT_TOKEN="Bot xxxxxxx"
```

## Development

To use a development version of Acropole, start by cloning the repository.

```
git clone https://gitlab.com/one-more/acropole.git
cd acropole
```

Then, you build it locally and run it with you own bot token.

```
scripts/build.dev.sh
scripts/run.dev.sh "Bot xxxxxxx"
```

### Testing

You can launch the test after building if you haven't done it yet.

```
scripts/build.dev.sh
```

Then, just run the test script.

```
scripts/test.dev.sh "Bot xxxxxxx"
```

