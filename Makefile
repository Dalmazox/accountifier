gen:
	buf generate
.PHONY: gen

mock:
	mockgen -source=internal/repositories/user_repository.go -destination=internal/repositories/mocks/mock_user_repository.go -package=mocks
	mockgen -source=internal/repositories/user_token_repository.go -destination=internal/repositories/mocks/mock_user_token_repository.go -package=mocks
	mockgen -source=internal/repositories/tx.go -destination=internal/repositories/mocks/mock_tx.go -package=mocks
.PHONY: mock