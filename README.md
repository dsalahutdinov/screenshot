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
docker run dsalahutdinov/screenshot -p 8080:8080
```

Than just curl for screenshot:

```sh
curl http://localhost:8080/screenshot/?width=1180&height=768&url=https://amplifr.com
```

## Kubernetes

To run on Kubernetes cluster:

```sh
kubectl create ns screenshot
kubectl create deployment -n screenshot --image=dsalahutdinov/screenshot screenshotkubectl create deployment -n screenshot --image=dsalahutdinov/screenshot screenshot
kubectl expose -n screenshot deployment screenshot --port=80 --target-port=8080
```

After than, screenshot is available on your cluster as the service:
```
curl http://screenshot.screenshot/screenshot/?width=1180&height=768&url=https://amplifr.com
```

