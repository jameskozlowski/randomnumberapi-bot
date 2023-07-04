
## About

This is a simple discord bot to generate random numbers using www.randomnumberapi.com

One of the features of this bot is to use reddit comments as seeds, to generate what I am calling a "free-will" seed.
The thought behind this, is that if you believe in free will, then this seed should be random and unable to predict

## Warnning
This API uses GO's random functions.
The Reddit random function does use the GO random function seeded by reddit comments.
This has not been tested for and should not be used for secure cryptographic purposes

## Discord bot token

You must provide a bot token by setting the BOT_TOKEN environment variable
```
export BOT_TOKEN="<token>"
```
## Build and run this using Podman/Docker

### Build
From the root of the project
```
sudo podman build -f ./docker/Dockerfile . -t radmonnumberapidiscordbot
```

### Run
```
sudo podman run -d --rm -e BOT_TOKEN='<token>' radmonnumberapidiscordbot
```

## Dependencies

| Name | Version | Disc                    |
|------|---------|-------------------------|
|github.com/bwmarrin/discordgo| v0.27.1 | discord bot lib         |
|github.com/gorilla/websocket| v1.4.2  | used by discord bot lib |
|golang.org/x/crypto| v0.1.0  | used by discord bot lib         |
|golang.org/x/sys| v0.1.0  | used by discord bot lib         |