PACKAGE_NAME=go-firebase-auth-server
SERVER=go-firebase-auth-server_server_1
LIMIT=0

.PHONY: migrate-status
migrate-status:
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate status"

.PHONY: migrate-new
migrate-new:
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate new $(NAME)"

.PHONY: migrate-up
migrate-up:
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate up"
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate status"

.PHONY: migrate-down
migrate-down:
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate down -limit=$(LIMIT)"
	docker exec -it $(SERVER) /bin/bash -c "sql-migrate status"

.PHONY: run
run:
	docker-compose up

.PHONY: build 
build:
	docker build -f build/prod/server/Dockerfile -t $(PACKAGE_NAME) .

.PHONY: submit
submit:
	docker tag $(PACKAGE_NAME):latest gcr.io/${PROJECT_ID}/$(PACKAGE_NAME):latest
	docker push gcr.io/${PROJECT_ID}/$(PACKAGE_NAME):latest


