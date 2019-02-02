# Forge
Autodesk API Client for golang.

## Setup
You'll need to [create an
app](https://forge.autodesk.com/en/docs/oauth/v2/tutorials/create-app/) in
order to use this client package. And then generate a `client id` and `client
secret`.

Once you have these keys, put them in a `.env` at the root of this project.

```
export FORGE_CLIENT_ID=XXXXXXXX
export FORGE_CLIENT_SECRET=XXXXXXXXX
```

## Development

#### Run tests

```sh
go test
```

## TODO
 * Setup a mock Autodesk API for testing instead of making real calls.
