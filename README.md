# GETDataConvertService
Simple Go Service which can read data in YAML, perform some calculations and transform it to JSON and send it back


Hello everyone! This is the latest version of my Go service. I strongly advice you to run this in container,
which is pushed to DockerHub. All you need to do is to run the following command:
```
docker run --network="host" -e ENDPOINT="your DB service name, which is basically server which will send you YAML"
nikitasadok/go-get-service
```
