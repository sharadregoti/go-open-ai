#!/bin/bash

my_var="
# What does the below Kubernetes resource YAML do?  Give detailed answers for each of the below questions & Ensure formatting.

What is this resource?
Why it is required?
How it does the intended work?

apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  annotations:
	kubectl.kubernetes.io/last-applied-configuration: |
	  {"apiVersion":"networking.istio.io/v1alpha3","kind":"Gateway","metadata":{"annotations":{},"name":"bookinfo-gateway","namespace":"jager-test"},"spec":{"selector":{"istio":"ingressgateway"},"servers":[{"hosts":["*"],"port":{"name":"http","number":80,"protocol":"HTTP"}}]}}
  creationTimestamp: "2023-02-17T11:12:56Z"
  generation: 1
  name: bookinfo-gateway
  namespace: jager-test
  resourceVersion: "38819715"
  uid: 3a807bd1-ede4-43c8-856a-86c8b9464368
spec:
  selector:
	istio: ingressgateway
  servers:
  - hosts:
	- '*'
	port:
	  name: http
	  number: 80
	  protocol: HTTP"

echo $my_var | ./go-open-ai
