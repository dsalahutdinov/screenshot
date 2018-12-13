# Browser screenshot service

Written with Golang, it uses the headless Google Chrome in Docker container to generate screenshot.

## Development

Run it locally: 
```sh
docker-compose up app
```

Navigate to the localhost:8080 in browser or with the `wget` tool:
```
wget http://localhost:8080/screenshot?url=https://amplifr.com
```
