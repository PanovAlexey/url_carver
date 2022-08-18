up:
	docker-compose --file .docker/docker-compose.yml up --build

down:
	docker-compose --file .docker/docker-compose.yml down

restart:
	docker-compose --file .docker/docker-compose.yml restart

recreate:
	docker-compose --file .docker/docker-compose.yml down
	docker-compose --file .docker/docker-compose.yml up -d --build --force-recreate
	docker-compose --file .docker/docker-compose.yml up -d

logs:
	cd .docker && docker logs go-template

build:
	go build -o ./build/shortener ./cmd/shortener

exec:
	cd .docker && docker-compose exec go-template bash

test:
	go test ./...

lint:
	golangci-lint run
