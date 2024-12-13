# redis-sentinel-client

Redis Sentinel client to simulate failover of Sentinel setup

Utilizing:
- Retries + Backoff
- `NewFailoverClusterClient` to route readonly commands to slave nodes
