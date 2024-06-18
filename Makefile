up-services:
	docker-compose up --build -d

down-services:
	docker-compose down

# Go commands
run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go /build

upd:
	go mod tidy