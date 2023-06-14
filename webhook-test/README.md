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

Webhook server allows:
```
kubectl apply -f config/test-deploy.yaml

# change config/test-deploy.yaml

kubectl apply -f config/test-deploy.yaml
```

Webhook server disallows:
```
# Edit config/webhook-server.yaml to use NOT_ALLOWED env

kubectl apply -f config/test-deploy.yaml
```
