# CF Example Golang Docker

## Run locally

```sh
# Build it
docker build -t jcscottiii/cf-ex-go-docker .
# Run it
docker run -it --rm -p 8000:8000 --name my-run-go jcscottiii/cf-ex-go-docker
```

# Push to CF
```
docker login
docker push jcscottiii/cf-ex-go-docker
cf push james-ex-go-docker --docker-image jcscottiii/cf-ex-go-docker -m 64M
```