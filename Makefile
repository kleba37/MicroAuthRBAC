run:
	go run ./cmd/main/main.go -mod=vendor

build:
	go build ./cmd/main/main.go -mod=vendor

tests:
	go test ./tests/*/*

migrate:
	go run ./cmd/Migrate/main.go -mod=vendor
