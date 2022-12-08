update-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	mkdir -p mocks/storage
	mockgen -source=internal/storage/storage.go > mocks/storage/storage.go