IMAGE_NAME = custom-prometheus
CONTAINER_NAME = prometheus
VOLUME_NAME = prometheus-data

#create_volume:
#	docker volume create \
#		--opt type=ext4 \
#		--opt device=/dev/sdb3 \
#		$(VOLUME_NAME)

# Build the Docker image
build_prometheus:
	docker build -t $(IMAGE_NAME) .


run: build_prometheus
	docker run -d \
		--name $(CONTAINER_NAME) \
		-v $(VOLUME_NAME):/prometheus \
		-p 9090:9090 \
		--network prom-net \
		$(IMAGE_NAME)

clean:
	docker rm -f $(CONTAINER_NAME) || true
	docker volume rm $(VOLUME_NAME) || true
	docker volume prune
	docker image rm -f $(IMAGE_NAME)
	
