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

.PHONY: test
test:
	go test ./internal/usecase -count=1 -v && go test ./internal/usecase -coverprofile=cover.txt && go tool cover -html cover.txt -o index.html

.PHONY: docker_pull
docker_pull:
	docker pull --platform linux/x86_64 prerec/exchanger:main

.PHONY: dockerhub_run
dockerhub_run:
	docker run -d -p 8080:8080 --platform linux/x86_64 prerec/exchanger:main