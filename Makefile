BACKEND_CONTAINER=todo-app
SQL_CONTAINER=todo-db

DOCKER_COMPOSE=$(shell echo 'docker-compose')

start: # Starts the containers
	$(DOCKER_COMPOSE) up -d
stop: # Stop the containers
	$(DOCKER_COMPOSE) down -v --remove-orphans
restart: # Restart the containers
	$(DOCKER_COMPOSE) restart $(BACKEND_CONTAINER)
app-logs: # Backend app logs
	$(DOCKER_COMPOSE) logs -f --tail 100 $(BACKEND_CONTAINER)
lint: # Go linting
	golint
	go vet ./...
	go fmt ./...
generate: # Generate gqlgen code
	docker exec -it $(BACKEND_CONTAINER) sh -c 'go generate ./...'
db-terminal:
	docker exec -it $(SQL_CONTAINER) psql -d todo todo1
rungql: # Runs the gqlgen
	docker exec -it $(BACKEND_CONTAINER) sh -c 'go run ./scripts/gqlgen.go'
