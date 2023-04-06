DOCKER_TAG=zhaohuabing/sample-controller:controller-runtime

build:
	env GOOS=linux GOARCH=amd64 go build ./controller.go
docker-build: build
	docker build -t $(DOCKER_TAG) ./
docker-push: docker-build
	docker push $(DOCKER_TAG)
deploy:
	kubectl create configmap sample-controller-webhook-crt --from-file=./tls.crt --from-file=./tls.key
	kubectl apply -f ./deploy.yaml
	kubectl apply -f ./webhook.yaml
.PHONY: build docker-build docker-push
