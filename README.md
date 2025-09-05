# Microservice Dashboard
A GraphQL-based microservice monitoring dashboard built with Go, GraphQL, and Kubernetes. This project simulates a microservicce architecture with dynamic metrics for each service. 

## Features
* GraphQL API with queries for all services or a single service by name.
* Simulated for multiple microservices (`auth`,`payments`,`checkout`,`inventory`,`notifications`).
* Ready for Kubernetes deployment.

## Tech Stack
* Go: backend
* GraphQL: API layer (gqlgen)
* Docker
* Kubernetes

## GraphQL Queries
1. `GetService`
```
  services {
    name
    status
    latencyMs
    errorRate
    cpuUsage
    memoryUsage
    uptime
    requestRate
  }
```
2. `GetServiceByName`
```
service(name: "auth") {
    name
    status
    latencyMs
    errorRate
    cpuUsage
    memoryUsage
    uptime
    requestRate
  }
```

## Running locally
1. Clone the repo:
```
git clone https://github.com/huhnbp/microservice-dashboard.git
cd microservice-dashboard
```
2. Install dependencies:
```
go mod tidy
```
3. Run the server:
```
go run server.go
```
4. Open [http://localhost:8080/](http://localhost:8080/)

## Docker
```
docker build -t microservice-dashboard .
docker run -p 8080:8080 microservice-dashboard
```

## Kubernetes
1. Make sure to have minikube or kind installed.
2. Apply the deployment:
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
3. Access the service:
```
kubectl port-forward svc/microservice-dashboard 8080:8080
```
4. Open [http://localhost:8080/](http://localhost:8080/)
