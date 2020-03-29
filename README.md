# middleware-example
An simple implementation of rate limiter using [go-redis](https://github.com/go-redis/redis_rate).  
It's also a show case of middleware usage in [gin](https://github.com/gin-gonic/gin).

## Quick start
1.  Clone and enter the repository.
```sh
    git clone https://github.com/sou2cute/middleware-example.git
    cd middleware-example
```
2.  Using docker-compose to start and attach to containers for services.
```sh
    docker-compose up
```
3.  Check ip location that docker machine is configured.
```sh
    docker-machine ip
```
See console, and you can ping the server now.
