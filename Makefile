build:
	cd cmd && go build -o bin/receipts_api

run: build
	export MYSQL_USER=root
	export MYSQL_PASSWORD=password
	export MYSQL_HOST=127.0.0.1:3306
	export MYSQL_DATABASE=db
	./cmd/bin/receipts_api

test:
	@go test -v ./...

update-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	mkdir -p mocks/storage
	mockgen -source=internal/storage/storage.go > mocks/storage/storage.go

