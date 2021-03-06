# sfapi [![Build Status](https://travis-ci.com/loubard/sfapi.svg?branch=master)](https://travis-ci.com/loubard/sfapi)

## Requirements

- Go
- sqlite

### Libraries

- gorm
- gorilla/mux

## Run

Running the project serves the api on port `8080`
```
go run
```

To test the URLs:

```bash
$ curl http://127.0.0.1:8080/v1/payments/  # list all payments
$ curl http://127.0.0.1:8080/v1/payments/{id} # get payment resource by id
$ curl -XDELETE http://127.0.0.1:8080/v1/payments/{id} # delete payment resource by id
$ curl -XPOST -d '{"payment_id":"42"}' http://127.0.0.1:8080/v1/payments/ # create payment resource
$ curl -XPUT -d '{"payment_id":"42"}' http://127.0.0.1:8080/v1/payments/{id} # update payment resource by id
```

where {id} is a payment id

## Test

```go
go run test /...

```

## To improve

- authentication
- validation
- godoc
- logging/monitoring
- deployment