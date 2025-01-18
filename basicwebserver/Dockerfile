FROM golang:1.23

WORKDIR /usr/src/app

EXPOSE 3333

COPY . .
RUN go mod download && go mod verify
RUN make build 

CMD ["./app"]