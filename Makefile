genType:
	gqlgen generate --config ./graphql/gqlgen.yml

bl-local:
	cd ./build && docker compose build --no-cache

up-local:
	cd ./build && docker compose up

inBackend:
	docker exec -it grahql-backend bash

InDb: 
	docker exec -it mysql-db bash