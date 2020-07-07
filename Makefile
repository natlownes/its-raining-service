IMAGE := narf/its-raining-service
PORT := 8080
ECR_NAME := its-raining-various-things-webservice:latest
ECR_URI := 016013605721.dkr.ecr.us-east-1.amazonaws.com/$(ECR_NAME)
REGISTRY = "10.30.2.110:32000"

app: clean
	@docker run --rm -v "$(CURDIR)":/opt/app -w /opt/app golang:1.8 go build

.PHONY:
image: app
	@docker build -t $(IMAGE) .

.PHONY:
push:
	docker tag $(IMAGE):latest $(REGISTRY)/$(IMAGE):latest
	docker push $(REGISTRY)/$(IMAGE):latest

.PHONY:
run: image
	docker run --rm -ti -p $(PORT):$(PORT) $(IMAGE)

.PHONY:
shell: image
	docker run --rm -ti -p $(PORT):$(PORT) $(IMAGE) --entrypoint="/bin/bash"

.PHONY:
clean:
	@rm -rf app

.PHONY:
login:
	@aws ecr get-login --no-include-email --region us-east-1 | bash

.PHONY:
upload: image login
	@docker tag $(IMAGE) $(ECR_URI)
	@docker push $(ECR_URI)
