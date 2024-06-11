# simple-rate-limiter
Simple rate limiter written in GO


# Run 
To run the server build the containers with

```sh
docker compose build
```

and run the golang app

```sh
docker exec simple-rate-limiter-service-1 go run main.go
```