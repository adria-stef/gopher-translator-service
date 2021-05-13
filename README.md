# gopher-translator-service :page_with_curl:

~

[![Build Status](https://travis-ci.org/adria-stef/gopher-translator-service.svg?branch=main)](https://travis-ci.org/adria-stef/gopher-translator-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/adria-stef/gopher-translator-service/goreportcard)](https://goreportcard.com/report/github.com/adria-stef/gopher-translator-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/adria-stef/gopher-translator-service/blob/main/LICENSE)

## Overview

This is an HTTP gopher translation service.

Gophers are friendly creatures but it’s not that easy to communicate with them. They have their own language and they don’t understand English. You can use this service to translate to their language.

## Setup and run locally

```sh
go mod vendor
go run main.go -port=8081
```

## Run tests

### Prerequisites

* [ginkgo](http://onsi.github.io/ginkgo/)

```sh
go generate ./...
ginkgo -v
```

## Sample calls locally

```sh
curl -X POST http://127.0.0.1:8080/word -d '{"english-word": "penguin"}'
```

```sh
curl -X POST http://127.0.0.1:8080/sentence -d '{"english-sentence": "You either die a hero, or you live long enough to see yourself become the villain."}'
```

```sh
curl http://127.0.0.1:8080/history
```
