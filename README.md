# Acropole

A Discord bot project to establish a (true) democracy.

## Getting Started

These instructions will get you a local discord bot running.

### Prerequisites

You just have to have Docker installed

```
docker run registry.gitlab.com/one-more/acropole:master -e BOT_TOKEN="Bot xxxxxxx"
```

### Development

This will launch bot with hot reloading

```
scripts/build.dev.sh
scripts/run.dev.sh "Bot xxxxxxx"
```

Launching tests:

```
scripts/test.dev.sh "Bot xxxxxxx"
```

