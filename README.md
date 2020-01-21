# Challenge

This repository is meant to do: 

* Implement a piece of software exposing a JSON documen when visited with a HTTP client
* Dockerizing the application
* MiniKubing it
* Implements second App utilizing the first and displays reversed message string
* Automating the deployment with a script
* ensuring running multiple instances of the application
* organize regular application upgrades

It is initially forked from the [cicdexample](https://github.com/cishiv/cicdexample) example. Making use of the build script which will automate the building of the first application, dockerizing it and putting it in minikube.

I followed another Approach using docker-compose as an automation tool for deploying the applications. For more info. about it and dependencies & Installation requirements could be found under: [Microservice challenge](https://github.com/Kareemabdallah/Microservice_Challenge)

## Rolling upgrade && Scaling out

we can use Kubernetes Rolling updates for rolling upgrades performed with zero-downtime by incrementally updating Pods instances with new ones. first by defining a yaml file such as the format below: 

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
