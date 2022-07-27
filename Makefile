run:
	docker-compose --env-file .env up
dep:
	go mod download
vet:
	go vet ./...
