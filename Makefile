
.PHONY: migrate-status
migrate-status:
	@echo run: sql-migrate status
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate status"

.PHONY: migrate-new
migrate-new:
	@echo run: sql-migrate new $(name)
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate new $(name)"

.PHONY: migrate-up
migrate-up:
	@echo run: sql-migrate up
	docker exec -it go-firebase-auth-server_server_1 /bin/bash -c "sql-migrate up"
