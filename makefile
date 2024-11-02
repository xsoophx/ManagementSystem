DOCKER_COMPOSE_FILE=docker-compose.yml
IMAGE_NAME=management-system
TAG=latest

include .env
export $(shell sed 's/=.*//' .env)

build:
	docker build -t $(IMAGE_NAME):$(TAG) .

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d

start:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi all --volumes --remove-orphans
