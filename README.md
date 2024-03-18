# Squid 🦑

A load-testing server written in Go meant to simulate concurrent users and requests per second while developing your applications.

It will offer up metadata about the health of your API once it's finished.

## Table of Contents

- [Getting Started](#getting-started)
- [Unit Testing](#unit-testing)
- [Contributing](#contributing)

## Getting Started

There is a Makefile for the project that contains the following commands you can use to interact with the load testing server:

```bash
# For building the Go binary
make build

# For running the server locally
make run

# For cleaning the build artifacts
make clean
```

Note that `make run` will prompt you to input a variety of data, including:

- Endpoint to load test
- Number of concurrent users to simulate

## Unit Testing

TODO

## Contributing

TODO

## Acknowledgements

TODO
