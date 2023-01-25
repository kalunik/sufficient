DB_PTH      = ./pgdata

all: build up

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

start:
	docker compose start

stop:
	docker compose stop

exec:
	docker exec -it $(serv) sh

rmvolume:
	rm -rf $(DB_PTH)
	docker volume prune -f

fclean: rmvolume down
	sudo docker system prune -af

re: down rmvolume build up

.PHONY: all build up down start stop rmvolume fclean re