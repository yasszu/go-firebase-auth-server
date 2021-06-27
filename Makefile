LIMIT=0

.PHONY: migrate-status
migrate-status:
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate status"

.PHONY: migrate-new
migrate-new:
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate new $(NAME)"

.PHONY: migrate-up
migrate-up:
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate up"
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate status"

.PHONY: migrate-down
migrate-down:
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate down -limit=$(LIMIT)"
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate status"
