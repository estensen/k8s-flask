# k8s-flask

Run locally
```
$ make up
```

Query the local app
```
$ curl -i http://localhost:5000/api/v1.0/books
```

Deploy to Kubernetes
```
kubectl apply -f manifest.yaml
kubectl expose deployment/k8s-flask-pod --type=NodePort --port 5000
```

Query the deployed app
```
$ curl -i http://localhost:31234/api/v1.0/books
```

Scale deployment
```
$ kubectl scale deployments/k8s-flask-pod --replicas=4
```

Cleaning up
```
$ kubectl delete services
$ kubectl delete deployment
```
