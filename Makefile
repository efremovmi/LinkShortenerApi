build:
	go build main.go
.PHONY: build

test:
	go test -v -race -timeout 30s ./...
.PHONY: test

create_migr:
	migrate create -ext sql -dir db/migrations -seq tabl_urls


migrate_up:
	sudo migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/Urls?sslmode=disable" -verbose up


migrate_down:
	sudo migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/Urls?sslmode=disable" -verbose down


show_dockers:
	sudo docker ps


show_pid:
	sudo lsof -i tcp:5432


stop_psql:
	sudo docker stop postgres12


start_psql:
	sudo docker start postgres12


drop_docker_d:
	sudo docker container rm postgres12


create_docker_db:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


create_db:
	sudo docker exec -it postgres12 createdb --username=root --owner=root Urls


drop_db:
	sudo docker exec -it postgres12 dropdb Urls

create_docker_images:
	sudo docker rm shortener_link_v1
	sudo docker rmi max/shortener_link:v1
	sudo docker build -t max/shortener_link:v1 --build-arg FLAG_STORE_WITH_DATABASE=false .
	sudo docker run -p 8000:8000 -d -e "FLAG_STORE_WITH_DATABASE=false" --name shortener_link_v1 max/shortener_link:v1


create_docker_images_compose:
	sudo docker stop linkshortenerapi_db_1
	sudo docker rm linkshortenerapi_db_1
	sudo docker rm linkshortenerapi_link_shortener_api_1
	sudo docker rmi linkshortenerapi_link_shortener_api
	sudo docker-compose up --build link_shortener