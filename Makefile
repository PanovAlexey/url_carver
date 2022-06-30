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
	docker container logs url_carver_go

build:
	go build -o ./build/shortener ./cmd/shortener

exec:
	docker exec -it url_carver_go sh

test:
	go test ./...

lint:
	golangci-lint run
