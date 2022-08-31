# 🚀 Spaceship Operator

This defines a `Spaceship` [CustomResourceDefinition](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) and a spaceship [Controller](https://kubernetes.io/docs/concepts/architecture/controller/).

The goal is to ensure the safety of all pods by evacuating the premises upon notice of blastoff. These pods will be rescheduled to return to their nodes following the launch of the spacecraft

## Usage

```bash
# Setup
kubectl create namespace spaceship
kubectl config set-context --current --namespace=spaceship

# Create CRD
kubectl apply -f spaceship-crd.yaml
kubectl api-resources --api-group=thomasvn.dev
kubectl explain spaceship

# Create objects
kubectl apply -f spaceship.yaml
kubectl get spaceships
kubectl describe spaceship nasa2022

# Build controller
go mod init spaceship
go mod tidy
go run main.go
```

## References

- <https://iximiuz.com/en/posts/kubernetes-operator-pattern>
- <https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/>
- <https://github.com/kubernetes/client-go/>
- <https://sdk.operatorframework.io/>

<!-- 
https://iximiuz.com/en/categories/?category=Kubernetes
-->

<!--
TODO:
- spaceship business logic in go?
  - https://pkg.go.dev/k8s.io/client-go@v0.25.0/kubernetes
  - How do people typically use the k8s client to monitor the CRD? Or do people make simple http requests?
  - Use the k8s client to figure out how to monitor the CRD
  - evict all pods for "blast off"
- Try using "kube-builder" (k8s sig) or "operator-sdk" (coreos)
- https://github.com/kubernetes/sample-controller
- Understanding Golang interfaces?
-->

<!--
DONE:
- Create Spaceship CRD
-->