Saga CLI
========

The [saga](./cmd/saga) CLI tool enables you to call the Saga APIs from the command line of your local development machine.

```
$ saga

Usage:
  saga [command]

Available Commands:
  booking     freight booking service
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
      --address string   address to connect to
  -h, --help             help for saga
      --insecure         make insecure client connection (only on localhost)
      --prod             connect to prod host (api.saga.einride.tech)
      --token string     bearer token used by client
  -v, --verbose          print verbose output

Use "saga [command] --help" for more information about a command.
```
