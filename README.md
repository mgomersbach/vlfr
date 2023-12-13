# Vault log filter and router

Listens on socket to process vault logs sent to it.
Can be used to filter out noise, or take action on certain patterns.
For example, notify a slack channel upon root token change request.
Or filter out sensitive information.

## Test it

## Build it

## Run it

### Configuration

vlfr takes a configuration file by default from /etc/vault/vlfr.yml

### Monitoring and metrics

Offers a /metrics endpoint for prometheus to scrape about the internals and statistics
