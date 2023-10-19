start-dev:
	docker-compose up --build
start-prod:
	docker-compose -f docker-compose-prod.yml up --build
test:
	cd backend && IS_TESTING=1 go test ./...