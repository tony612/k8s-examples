## Build image

```
docker build . -t webhook-server
```

## Deploy

```
kubectl apply -k config/
kubectl -n webhook-test rollout restart deploy webhook-server
```

## Verify

```
kubectl apply -f config/test-deploy.yaml
```
