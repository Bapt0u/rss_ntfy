# RSS Notify
[![Docker](https://github.com/Bapt0u/rss_ntfy/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/Bapt0u/rss_ntfy/actions/workflows/docker-publish.yml)

Basic RSS notifier. Read specific feeds and forward them to a specific webhook.

## `./conf/feeder.yml`

```yml
notify: https://discord_or_whatever

feeds:
  - https://github.com/hashicorp/vault/releases.atom
  - https://github.com/hashicorp/...t/releases.atom
```
## Build from sources

```bash
git clone git@github.com:Bapt0u/rss_notify.git rss_ntfy
cd rss_ntfy
go build -o rss_ntfy main.go
./rss_ntfy

# Build docker image
docker buildx build -t registry.dev.localhost/rss_ntfy:v1.0
docker run -it --rm -name rss_ntfy -v $(pwd)/conf registry.dev.localhost/rss_ntfy:v1.0

# Let's goooo
```

## ROADMAP

- [x] Matrix formated messages
- [ ] Define other message format (ntfy.sh, discord...)
- [ ] Add exception to notification based on the tag. If it contains a certain pattern, ignore it. (E.g. [Vault api version uprade](https://github.com/hashicorp/vault/releases/tag/api%2Fv1.12.0) I want to ignore this pattern, as I only care about Vault itself.)
