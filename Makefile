up-services:
	docker-compose up --build -d

down-services:
	docker-compose down

reset-db:
	make down-services
	rm -rf dbdata
	make up-services

# Go commands
run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go /build

upd:
	go mod tidy