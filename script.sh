#!/bin/bash
source ~/.bashrc

echo "Building"

# Build docker image
docker build -t test/api:latest ./api

# Build a docker image
buildDocker() {
	echo "🐋	$(timestamp): building image example:test"
	docker build -t example:test .
}

# Deploy to Minikube using kubectl
deploy() {
	echo "🌧️	 $(timestamp): deploying to Minikube"
	kubectl delete deployment example
	kubectl delete service example
	kubectl apply -f deploy.yml
}

# Run new one
docker run --name test-api -d -p 8080:2000 \
		-e ENV=production \
		-e DB_HOST=mongodb://admin:pass@datastore:27017 \
		-e DB_NAME=irs \
		test/api:latest
