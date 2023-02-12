sqlc:
	./scripts/sqlc.sh

build:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o main ./cmd/main.go
