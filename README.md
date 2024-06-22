# Satoshi Exchange API

The Sample of Crypto Exchange

## Getting Started

```bash
$ cp .env.example .env # Only at the beginning.
$ make up
```

## Feature

- Authentication(Firebase)
- REST API
- DB Transaction
- Auto Migration
- Unit Test
- Integration Test(Not all APIs, but some)

### Unit Test

`TEST_PATH` is paths under /tests/unit

example
```bash
$ make unit_test TEST_PATH=/lib/common_test.go
```

### Integration Test

`TEST_PATH` is paths under /tests/integration    
It is necessary to launch docker compose for testing in advance.

example
```bash
$ make up_test # launch docker compose for testing
$ make integration_test TEST_PATH=/controllers/trade_controller_test.go
```
