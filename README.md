# Challenge

This repository is meant to do: 

* Implements an app written in go exposing a JSON document when visited with a HTTP client
* Dockerizes the application.
* MiniKubing it.
* Implements second App utilizing the first and displays reversed message string.
* Automates the deployment with a script.
* Ensures running multiple instances of the application.
* Organizes regular application upgrades.

```Build.sh ``` automates the building of the first application, dockerizing it and putting it in minikube.

I followed another Approach using docker-compose as an automation tool for deploying the applications. For more info. about it and dependencies & Installation requirements could be found under: [Microservice challenge](https://github.com/Kareemabdallah/Microservice_Challenge)

## Go application
For this Approach, build script is used to build the an Go App which serve HTTP server for the application under localhost:9000. The json file is stored in */static* folder. we can test the app with ```go build```. The second App will be listening on localhost:7000.

Using 2-stage Dockerfile, It will first create a builder image with the entire contents of the local directory copied into */build* directory, which will then produce a binary called *app1* *app2*.

I stored the *app1* *app2* into two different folders to avoid package redeclaration errors. 

## Containerisation

The second stage creates the image will run based on Alpine image. It will use */app* as home directory. Image can be tested by running ```docker build -t app1:test .```. Building the docker image command need to be run ```./build build docker```. Then we could Minikube the app using the command Using ```kubectl apply -f deploy.yml```

## Building

Building the applications with ```go build -o app1``` ```go build -o app2```

Finally builder could be run using ```./app1``` ```./app1```

## Automation

Running the build script for automating the building process and running of the docker images then the deployment of the *app1* deployment and service using deploy file.

```
./build.sh
```

## Rolling upgrade && Scaling out

One can use Kubernetes Rolling updates for rolling upgrades performed with zero-downtime by incrementally updating Pods instances with new ones. We can also specify the number of instances required for the particular App by defining *replicas* to e.g. 2. First by defining a yaml file such as the format below: 

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: micro
spec:
  selector:
    matchLabels:
      name: micro
  replicas: 2
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        name: app1
    spec:
      containers:
        - name: micro
          image: rolling:v1
          imagePullPolicy: IfNotPresent
          ports:
            - name: micro
              containerPort: 9000
          resources:
            limits:
              cpu: "1"
              memory: 512Mi
            requests:
              cpu: "1"
              memory: 512Mi
          readinessProbe:
            httpGet:
              path: /health
              port: 9000
            initialDelaySeconds: 5
            periodSeconds: 2
            failureThreshold: 20
---
apiVersion: v1
kind: Service
metadata:
  name: micro
  labels:
    name: app1
spec:
  type: NodePort
  ports:
    - port: 9000
      #nodePort: 8888
  selector:
    name: app1
```
