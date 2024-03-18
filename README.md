# Squid 🦑

A load-testing server written in Go meant to simulate concurrent users and requests per second while developing your applications.

It will offer up metadata about the health of your API once it's finished.

## Table of Contents

- [Getting Started](#getting-started)
- [Unit Testing](#unit-testing)
- [Project Management](#project-management)
- [Contributing](#contributing)
- [Acknowledgements](#acknowledgements)

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
- Length of time in seconds to ramp to maximum concurrent users
- Length of time in seconds to run the load test

## Unit Testing

TODO

## Project Management

You can find a list of planned changes (past, present, future) on [this Trello board](https://trello.com/invite/b/2WRhAcKG/ATTI423f67efa34216844f95691652bb1d406B165047/load-testing-server).

## Contributing

TODO

## Acknowledgements

TODO
