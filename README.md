# redis-sentinel-client

Redis Sentinel client to simulate failover of Sentinel setup

Utilizing:
- Retries + Backoff
- `NewFailoverClusterClient` to route readonly commands to slave nodes


## Usage
k run sample-redis-client-retry --image=hudani/redis-sentinel-client:6