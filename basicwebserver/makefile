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
	