FROM golang:1.23

WORKDIR /usr/src/app

EXPOSE 3333

# These build args are passed by buildx automatically
ARG TARGETARCH

COPY . .
RUN go mod download && go mod verify
#RUN make build 
RUN if [ "$TARGETARCH" = "arm64" ]; then \
      make build-arm64 && cp app_arm app ; \
    else \
      make build-amd64 && cp app_amd64 app ; \
    fi

CMD ["./app"]