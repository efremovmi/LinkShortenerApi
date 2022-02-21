build:
	go build main.go
.PHONY: build

test:
	go test -v -race -timeout 30s ./...
.PHONY: test

create_migr:
	migrate create -ext sql -dir db/migrations -seq tabl_urls


migrate_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/Urls?sslmode=disable" -verbose up


migrate_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/Urls?sslmode=disable" -verbose down


show_dockers:
	sudo docker ps

show_pid:
	sudo lsof -i tcp:5432


stop_psql:
	sudo docker stop postgres12


start_psql:
	sudo docker start postgres12


drop_docker:
	sudo docker container rm postgres12


create_docker:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root Urls


dropdb:
	sudo docker exec -it postgres12 dropdb Urls