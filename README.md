# Pre-requisites
- Install Docker
- Install Docker Compose

# Run
```
docker-compose up -d
```

# Check

You can check the health of the service via http://IP:8888/health

# Test

You can now access the service via http://IP:8888/hello?name=abc

# Cleaning up

```
docker-compose down
```