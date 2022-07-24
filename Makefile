BINARY_NAME=userservice

clean:
	rm -r build
build: clean
	go build -o build/${BINARY_NAME} ./cmd/api
run: build
	docker compose up
dep:
	go mod download
vet:
	go vet
