DOCKER_TAG=zhaohuabing/demo-controller:leader-election

build:
	env GOOS=linux GOARCH=amd64 go build ./controller.go
docker-build: build
	docker build -t $(DOCKER_TAG) ./
docker-push: docker-build
	docker push $(DOCKER_TAG)
deploy:
	kubectl apply -f ./deploy.yaml
.PHONY: build docker-build docker-push
