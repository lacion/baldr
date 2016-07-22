.PHONY: build

IMAGE_NAME := "kyperion/baldr"

default: build

build: 
	docker build -f Dockerfile.build -t $(IMAGE_NAME):build .
	docker run $(IMAGE_NAME):build cat /gopath/bin/baldr > baldr
	chmod +x baldr