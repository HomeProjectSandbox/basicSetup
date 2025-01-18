run:
	docker network create prom-net
	cd basicwebserver && make docker_run
	cd prometheus && make run
clean:
	cd basicwebserver && make clean
	cd prometheus && make clean
	docker network rm prom-net
	