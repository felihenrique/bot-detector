start-dev:
	docker-compose up --build
start-prod:
	docker-compose -f docker-compose-prod.yml up --build
test:
	cd backend && IS_TESTING=1 go test ./...
migrate:
	cd backend && go run migrations/migrations.go
locust-master:
	locust --web-host=localhost --master
locust-worker:
	locust --web-host=localhost --worker