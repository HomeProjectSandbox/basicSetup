IMAGE_NAME=my-golang-app
CONTAINER_NAME=my-running-app

run:
	go run .
build:
	go build -o app .
build_and_run:
	go build -o app .
	./app
docker_run:
	docker build -t $(IMAGE_NAME) .
	docker run -d -it --rm --name $(CONTAINER_NAME) -p 3333:3333 --network monitoring $(IMAGE_NAME)
clean:
	docker container rm -f $(CONTAINER_NAME)
	docker image rm -f $(IMAGE_NAME)

build_and_push:
	@if [ "$(shell uname -m)" = "arm64" ] || [ "$(shell uname -m)" = "aarch64" ]; then \
		docker build -t tamasujvari/basicsetup_basicwebserver:arm .; \
		docker push tamasujvari/basicsetup_basicwebserver:arm; \
	else \
		docker build -t tamasujvari/basicsetup_basicwebserver:latest .; \
		docker push tamasujvari/basicsetup_basicwebserver:latest; \
	fi
