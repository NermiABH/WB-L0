up:
	docker-compose up -d
rebuild:
	docker-compose up -d --no-deps --build service
stop:
	docker-compose stop
publisher:
	go run cmd/publisher/main.go
service:
	go run cmd/service/main.go