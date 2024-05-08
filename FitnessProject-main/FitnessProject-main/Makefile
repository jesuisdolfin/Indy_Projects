postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it fitnessdb createdb --username=root --owner=root fitnessdb

dropdb:
	docker exec -it fitnessdb dropdb fitnessdb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fitnessdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fitnessdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

#docker status for all containers
#docker ps -a

# #stop running a container
# docker stop (container name)

# #start an existing container
# docker start (container name)

# #pull a docker image
# docker pull (image):(tag)

# #start a new container
# docker run --name (container name) -e (environment variables) -p (hostpost):(containerport) -d (image):(tag)

# #run command in container
# docker exec -it (container name) (command)[args]

# #workaround for generating sqlc code from user.sql
# docker run --rm -v c/users/charl/goproject:/src -w /src kjconroy/sqlc generate

.PHONY: postgres12 createdb dropdb