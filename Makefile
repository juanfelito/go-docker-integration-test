test-integration:
	docker compose build example-db server integration-tests
	docker compose up -d example-db
	sleep 5
	docker compose up -d server
	sleep 5
	docker compose run integration-tests
	docker compose down