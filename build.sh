# !/bin/sh

# Timestamp Function
timestamp() {
	date +"%T"
}

# Temporary file for stderr redirects
tmpfile=$(mktemp)

# Go build
    echo "⏲️	$(timestamp): started build script..."
    echo "🐋	$(timestamp): building image app1:test"
    echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
    docker build -t app1:test -f Dockerfile .
    docker images
    docker run -dti --rm -p 9000:9000 app1:test
    docker push $DOCKER_USERNAME/app2:test

    docker build -t app2:test -f Dockerfile2 .
    docker run -dti --rm -p 7000:7000 app2:test
    docker push $DOCKER_USERNAME/app2:test
    

echo "🌧️	 $(timestamp): deploying to Minikube"
	kubectl delete deployment app1
	kubectl delete service app1
	kubectl apply -f deploy.yml