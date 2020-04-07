<a href="https://amplifr.com/?utm_source=screenshot">
  <img width="100" height="140" align="right"
    alt="Sponsored by Amplifr" src="https://amplifr-direct.s3-eu-west-1.amazonaws.com/social_images/image/37b580d9-3668-4005-8d5a-137de3a3e77c.png" />
</a>
# Browser screenshot service

Written with Golang, it uses the headless Google Chrome in Docker container to generate screenshot.

## Development

Run it locally: 
```sh
docker-compose up app
```

Navigate to the localhost:8080 in browser or with the `wget` tool:
```
wget http://localhost:8080/screenshot?width=1180&height=768&url=https://amplifr.com
```

## Run with Docker

```sh
docker run -p 8080:8080 dsalahutdinov/screenshot
```

Than just curl for screenshot:

```sh
curl http://localhost:8080/screenshot/?width=1180&height=768&url=https://amplifr.com
```

## Helm

Install screenshot service from [Helm repository](https://github.com/dsalahutdinov/screenshot-helm):

```sh
$ helm repo add screenshot-helm https://dsalahutdinov.github.io/screenshot-helm/
"screenshot-helm" has been added to your repositories
$ helm install screenshot-helm/screenshot --name screenshot --namespace screenshot
NAME:   screenshot
LAST DEPLOYED: Sat Apr  4 16:59:02 2020
NAMESPACE: screenshot
STATUS: DEPLOYED

RESOURCES:
==> v1/Deployment
NAME         AGE
screenshot  1s

==> v1/Pod(related)
NAME                      AGE
screenshot-76948b-rrq77  1s

==> v1/Service
NAME         AGE
screenshot  1s
```

## Kubernetes

To run on Kubernetes cluster:

```sh
kubectl create ns screenshot
kubectl create deployment -n screenshot --image=dsalahutdinov/screenshot screenshot
kubectl expose -n screenshot deployment screenshot --port=80 --target-port=8080
```

After than, screenshot is available on your cluster as the service:
```
curl http://screenshot.screenshot/screenshot/?width=1180&height=768&url=https://amplifr.com
```

