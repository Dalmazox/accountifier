	DB_CONN ?= "postgres://docker:d0cker@localhost:5432/accountifier?sslmode=disable"

gen:
	buf generate
.PHONY: gen

mock:
	mockgen -source=internal/repositories/user_repository.go -destination=internal/repositories/mocks/mock_user_repository.go -package=mocks
	mockgen -source=internal/repositories/user_token_repository.go -destination=internal/repositories/mocks/mock_user_token_repository.go -package=mocks
	mockgen -source=internal/repositories/tx.go -destination=internal/repositories/mocks/mock_tx.go -package=mocks
.PHONY: mock

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
.PHONY: coverage

migrate-dev-up:
	migrate -path ./migrations -database $(DB_CONN) up
.PHONY: migrate-dev-up

migrate-dev-down:
	migrate -path ./migrations -database $(DB_CONN) down
.PHONY: migrate-dev-down
