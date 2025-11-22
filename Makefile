run:
	go run ./cmd/main/main.go -mod=vendor

build:
	go build ./cmd/main/main.go -mod=vendor

test:
	go test ./tests/*/**

migrate:
	go run ./cmd/migrate/main.go -mod=vendor
