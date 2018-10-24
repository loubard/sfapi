# sfapi [![Build Status](https://travis-ci.com/loubard/sfapi.svg?branch=master)](https://travis-ci.com/loubard/sfapi)

## Requirements

- Go

## Run

Running the project serves the api on port `8080`
```
go run
```

To test the URL:

```bash
$ curl http://127.0.0.1:8080/v1/payments/{id}
```

where {id} is a payment id

## Test

```go
go run test /...

```
