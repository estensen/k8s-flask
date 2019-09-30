# k8s-flask

## Run locally
```
$ make up
```

Query the local app:
```
$ curl http://localhost:5000/api/v1.0/books
[
  {
    "author": "Rene Redzepi,  David Zilber",
    "id": 0,
    "published": "2018",
    "title": "The Noma Guide to Fermentation"
  },
  {
    "author": " Donald A. Norman",
    "id": 1,
    "published": "2013",
    "title": "The Design of Everyday Things"
  },
  {
    "author": " Andreas M. Antonopoulos",
    "id": 2,
    "published": "2014",
    "title": "Mastering Bitcoin"
  }
]
```

### Access Jaeger UI
Open `http://localhost:16686/`


## Deploy to Kubernetes
Push image to Docker Hub
```
$ docker push estensen/books
```

Create deployment
```
$ kubectl apply -f manifest.yaml
```

Query the deployed app:
```
$ curl http://localhost:31234/api/v1.0/books
[
  {
    "author": "Rene Redzepi,  David Zilber",
    "id": 0,
    "title": "The Noma Guide to Fermentation",
    "year_published": "2018"
  },
  {
    "author": " Donald A. Norman",
    "id": 1,
    "published": "2013",
    "title": "The Design of Everyday Things"
  },
  {
    "author": " Andreas M. Antonopoulos",
    "id": 2,
    "published": "2014",
    "title": "Mastering Bitcoin"
  }
]
```

Scale deployment:
```
$ kubectl scale deployments/books --replicas=4
```

Cleaning up:
```
$ kubectl delete services books && kubectl delete deployment books
```
