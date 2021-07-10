PACKAGE_NAME=go-firebase-auth-server
SERVER=go-firebase-auth-server_server_1
LIMIT=0

.PHONY: migrate-status
migrate-status:
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate status -env="development"'

.PHONY: migrate-new
migrate-new:
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate new $(NAME) -env="development"'

.PHONY: migrate-up
migrate-up:
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate up -env="development"'
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate status -env="development"'

.PHONY: migrate-down
migrate-down:
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate down -env="development" -limit=$(LIMIT)'
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate status -env="development"'

.PHONY: migrate-up-test
migrate-up-test:
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate up -env="test"'
	docker exec -it $(SERVER) /bin/bash -c 'sql-migrate status -env="test"'

.PHONY: build
build:
	docker compose build

.PHONY: run
run:
	docker compose up

.PHONY: lint
lint:
	docker compose exec server golangci-lint run

.PHONY: test
test:
	docker compose exec server go test ./... -count=1

.PHONY: docker-build
docker-build:
	docker build -f build/prod/server/Dockerfile -t $(PACKAGE_NAME) .

.PHONY: docker-submit
docker-submit:
	docker tag $(PACKAGE_NAME):latest gcr.io/${PROJECT_ID}/$(PACKAGE_NAME):latest
	docker push gcr.io/${PROJECT_ID}/$(PACKAGE_NAME):latest
