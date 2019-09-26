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
$ kubectl apply -f manifest.yaml
```

Query the deployed app
```
$ curl -i http://localhost:31234/api/v1.0/books
```

Scale deployment
```
$ kubectl scale deployments/k8s-flask --replicas=4
```

Cleaning up
```
$ kubectl delete services k8s-flask
$ kubectl delete deployment k8s-flask
```
