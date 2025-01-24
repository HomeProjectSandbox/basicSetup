run:
	docker network create monitoring
	cd basicwebserver && make docker_run
	cd prometheus && make run
clean:
	cd basicwebserver && make clean
	cd prometheus && make clean
	docker network rm monitoring
	
docker_c_up:
	docker-compose up -d

docker_c_clean:
	docker-compose down --volumes --remove-orphans
	docker volume prune -f
	docker rmi -f $(docker images -aq)