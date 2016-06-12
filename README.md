# pp
PP is a Popular Purchases microservice for the Discount Ascii Warehouse ecommerce platform.

### Usage
```bash
# Upgrade your local version of the repo.
go get -u github.com/alkchr/pp

# Change location to the root of the project.
cd $GOPATH/src/github.com/alkchr/pp

# Run the application for testing.
go run main.go --purchases:api.uri="http://74.50.59.155:6000/api/"
```
Optionally, change parameters of the `config/app.ini` before running the app.

### Status
[![GoDoc](https://godoc.org/github.com/alkchr/pp?status.svg)](https://godoc.org/github.com/alkchr/pp)
[![Build Status](https://travis-ci.org/alkchr/pp.svg?branch=master)](https://travis-ci.org/alkchr/pp)
[![Coverage](https://codecov.io/github/alkchr/pp/coverage.svg?branch=master)](https://codecov.io/github/alkchr/pp?branch=master)
[![Go Report Card](http://goreportcard.com/badge/alkchr/pp?t=3)](http://goreportcard.com/report/alkchr/pp)

### License
Distributed under the terms of MIT license unless otherwise noted.
