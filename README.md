# Operator

Playing with CRDs & Controllers 🚀

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

<!-- 
https://iximiuz.com/en/categories/?category=Kubernetes
-->

<!--
TODO:
- spaceship business logic in go?
  - play around with the kubernetes client in go
  - https://pkg.go.dev/k8s.io/client-go@v0.25.0/kubernetes
  - Use the k8s client to figure out how to monitor the CRD
  - evict all pods for "blast off"
-->

<!--
DONE:
- Create Spaceship CRD
-->