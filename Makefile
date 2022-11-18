update-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	mkdir -p mocks/repositories
	mkdir -p mocks/services
	mockgen -source=internal/repositories/receipts_repository.go > mocks/repositories/receipts_repository_mocks.go
	mockgen -source=internal/services/receipts_service.go > mocks/services/receipts_service_mocks.go