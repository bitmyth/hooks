PWD := $(shell pwd)
IMAGE := bitmyth/hooks

image:
	docker build -t $(IMAGE) .

## Run Docker production image
run-prod:
	docker run --rm --name hooks -p 8000:8000 -e FILE=jobs.yaml $(IMAGE)

## Run on local
run:
	go run src/server/main.go

