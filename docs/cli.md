# Saga CLI

The [saga](./cmd/saga) CLI tool enables you to call the Saga APIs from the
command line of your local development machine.

This cli is installed on running `make`. You can set a default user by using
credentials of an Api connection from your organization:

```shell
saga auth login --client-id <CLIENT_ID> --client-secret <CLIENT_SECRET>
```

This will store a file with the credentials locally. On every request the cli
will fetch a fresh auth token, based on these credentials. After this you can
use the cli to query saga, like:

```shell
saga shipment get-shipment --name spaces/123/shipments/234 --us-prod
```

To clear the user credentials use:

```shell
saga auth logout
```

```
$ saga 

Saga CLI

USAGE
  saga <command>

SERVICE COMMANDS
  authentication           Authentication service
  booking                  Freight booking service
  shipment                 Shipment service
  shipment-tracking-event  Shipment tracking event service

MODULE COMMANDS
  auth  authentication commands for cli

OTHER COMMANDS
  help        Help about any command
  completion  Generate the autocompletion script for the specified shell

HOST FLAGS
  --eu-prod    bool  Connect to eu-prod host (eu.api.saga.einride.tech)
  --us-prod    bool  Connect to us-prod host (us.api.saga.einride.tech)

CONNECTION FLAGS
  --address     string  Connect to address
  --insecure    bool    Make insecure connection (only on localhost)
  --token       string  Authenticate with a bearer token

OTHER FLAGS
  -h  --help       bool  Show help for command
  -v  --verbose    bool  Enable verbose output

```
