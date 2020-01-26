#!/bin/bash
docker tag app1:latest mostafa/app1:version1.0.test
docker push mostafa/app1:version1.0.test
docker tag app2:latest mostafa/app2:version1.0.test
docker push mostafa/app2:version1.0.test