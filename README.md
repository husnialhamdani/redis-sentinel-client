# redis-sentinel-client

Redis Sentinel client to simulate failover of Sentinel setup

Utilizing:
- Retries + Backoff
- `NewFailoverClusterClient` to route readonly commands to slave nodes


## Usage
k run sample-redis-client-retry --image=hudani/redis-sentinel-client:6

## sample
![alt text](http://url/to/img.png)
![alt text](https://github.com/husnialhamdani/redis-sentinel-client/blob/main/redis-sentinel-failover-test.png?raw=true)