COMPOSE_PTH	= ./docker-compose.yml
DB_PTH      = ./pgdata

all: vol build up

build:
	docker compose -f $(COMPOSE_PTH) build

up:
	docker compose -f $(COMPOSE_PTH) up

ps:
	docker ps -a

logs:
	docker compose -f $(COMPOSE_PTH) logs


stop:
	docker compose -f $(COMPOSE_PTH) stop

start:
	docker compose -f $(COMPOSE_PTH) start

down:
	docker compose -f $(COMPOSE_PTH) down

exec:
	docker exec -it $(serv) sh

vol:
	mkdir -p $(DB_PTH)

rmvol:
	sudo rm -rf $(DB_PTH)
	sudo docker volume prune -f

fclean: rmvol down
	sudo docker system prune -af

re: down rmvol vol build up

.PHONY: all hosts build up ps logs stop start down vol rmvol fclean re