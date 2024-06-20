# Satoshi Exchange API

The Sample of Crypto Exchange

[Behavior Videos](./docs/behavior_videos)

## Setup

```
cp .env.example .env # Only at the beginning.
make up
```

## Feature

- Authentication(Firebase)
- REST API
- MVC Model
  - [src/internal](./src/internal)
- Error Handling
- DB Transaction
  - [src/internal/services/tx_service.go](./src/internal/services/tx_service.go)
- Unit Test
  - [src/internal/tests/unit](./src/internal/tests/unit)
- Integration Test(Not all APIs, but some)
  - [src/internal/tests/integration](./src/internal/tests/integration)

### Unit Test

TEST_PATH is paths under /tests/unit

example
```
make unit_test TEST_PATH=/lib/common_test.go
```

### Integration Test

TEST_PATH is paths under /tests/integration    
It is necessary to launch docker compose for testing in advance.

example
```
make up_test # launch docker compose for testing
make integration_test TEST_PATH=/controllers/trade_controller_test.go
```
