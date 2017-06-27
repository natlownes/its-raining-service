IMAGE := narf/its-raining-service

build:
	@docker run --rm -v "$(CURDIR)":/opt/app -w /opt/app golang:1.8 go build

image: build
	@docker build -t $(IMAGE) .

run: image
	docker run --rm -ti -p 8080:8080 $(IMAGE)

clean:
	@rm -rf app
