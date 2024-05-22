.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: docs
docs:
	swag init -g cmd/app/main.go

.PHONY: docker_build
docker_build:
	docker build . -t exchanger

.PHONY: docker_run
docker_run:
	docker run -d -p 8080:8080 exchanger