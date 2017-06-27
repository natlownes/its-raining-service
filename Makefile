IMAGE := narf/its-raining-service
PORT := 8080

build:
	@docker run --rm -v "$(CURDIR)":/opt/app -w /opt/app golang:1.8 go build

image: build
	@docker build -t $(IMAGE) .

run: image
	docker run --rm -ti -p $(PORT):$(PORT) $(IMAGE)

clean:
	@rm -rf app
