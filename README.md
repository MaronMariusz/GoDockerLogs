# GoDockerLogs
Simple CLI script for reading docker logs files by its timestamp

# How to use

1. Export docker logs with timestamp data `docker logs -t > logs.txt`
2. `go run read.go ./logs.txt`
3. Open `http://localhost:9090`
4. Select log line by timestamp

