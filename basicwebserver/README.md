`docker build -t my-golang-app .`

-p [HOST_PORT]:[CONTAINER_PORT]
`docker run -d -it --rm --name my-running-app -p 3333:3333 my-golang-app`