# 🍉 PassGen

Simple password generator.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes. See deployment for notes on how to deploy the project on a live system.

More documentation:

```shell
godoc -http=:6060
open http://localhost:6060/pkg/github.com/rjNemo/go-pass-gen/
```

### Prerequisites

- Install a recent Golang version on the [official website](https://golang.org/doc/install).
- Install `air` to enable auto reload during development using :

```shell
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

Then initialize it in the project directory with:

```shell
air init
```

### Installing

Install dependencies using:

```shell
go mod tidy
```

Run the development server with:

```shell
air
```

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```shell script
Give an example
```

### And coding style tests

Explain what these tests test and why

```shell script
Give an example
```

## Deployment

[Quality check](https://goreportcard.com/) before deployment.

## Built With

| Tool                                       | Description                                                                                                     |
| ------------------------------------------ | --------------------------------------------------------------------------------------------------------------- |
| [Go](https://golang.org/)                  | Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. |
| [Cobra](https://cobra.dev/) | A Framework for Modern CLI Apps in Go                                                                                          |

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see
the [tags on this repository](https://github.com/rjNemo/go-pass-gen/tags).

## Authors

- **Ruidy** - _Initial work_ - [Ruidy](https://github.com/rjNemo)

See also the list of [contributors](https://github.com/rjNemo/go-pass-gen/contributors) who participated in this
project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
