docker-build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose stop

docker-clean:
	docker-compose rm -f