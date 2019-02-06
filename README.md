[![Build Status](https://travis-ci.org/cbarraford/forge.svg?branch=master)](https://travis-ci.org/cbarraford/forge)
[![codecov](https://codecov.io/gh/cbarraford/forge/branch/master/graph/badge.svg)](https://codecov.io/gh/cbarraford/forge)
[![GoDoc](https://godoc.org/github.com/cbarraford/forge?status.svg)](https://godoc.org/github.com/cbarraford/forge)

# Forge
Autodesk API Client for golang.

## Status
This project is consider "early phase", and therefore may not implement all
aspects of the Forge API, is subject to change, and may not have detail tests
written.

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

Once, you've done that you can create a `Client` via...

```go
client, err := forge.New()
```


Alternatively, you pass your creds in your code and receive a `Client`.

```go
client, err := forge.NewWithCreds("IDXXXX", "SECRETXXXX")
```

## Development

#### Run tests

```sh
go test
```

## TODO
 * Setup a mock Autodesk API for testing instead of making real calls.
