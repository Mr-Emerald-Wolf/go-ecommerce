# Define variables
DOCKER_COMPOSE = docker-compose

# Targets
.PHONY: build up down logs restart clean

build:
	docker build . -t shivam2002/go-ecommerce:latest

up:
	$(DOCKER_COMPOSE) up -d --build 

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f

restart:
	$(DOCKER_COMPOSE) restart

clean:
	$(DOCKER_COMPOSE) down -v

# Help target
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build      Build Docker containers"
	@echo "  up         Start Docker containers in the background"
	@echo "  down       Stop and remove Docker containers"
	@echo "  logs       View logs of Docker containers"
	@echo "  restart    Restart Docker containers"
	@echo "  clean      Stop, remove containers, and also remove volumes (data)"
	@echo "  help       Display this help message"