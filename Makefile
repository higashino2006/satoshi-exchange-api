TEST_PATH=

up:
	docker compose up -d --build

up_test:
	docker-compose -f docker-compose.test.yml -p se-api-test up -d --build

down:
	docker compose down

down_test:
	docker compose down -p se-api-test

unit_test:
	go test "./src/internal/tests/unit$(TEST_PATH)"

integration_test:
	docker compose exec go -p se-api-test /bin/bash \
	go test "./src/internal/tests/integration$(TEST_PATH)"
