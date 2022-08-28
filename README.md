# Operator

Playing with CRDs & Controllers 🚀

## Setup

```bash
kubectl create namespace spaceship

kubectl apply -f spaceship-crd.yaml -n spaceship
kubectl api-resources --api-group=thomasvn.dev
kubectl explain spaceship
```

## References

- <https://iximiuz.com/en/posts/kubernetes-operator-pattern>

<!-- 
https://iximiuz.com/en/categories/?category=Kubernetes
-->

<!--
TODO:
- spaceship business logic in go?
-->