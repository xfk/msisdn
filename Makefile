dev:
	docker-compose down
	docker-compose up -d

test:
	go test
