# GinGormRedisApi


## Build
```bash
    docker build -t gin-gorm-redis-api .
```

## Run
```bash
    sudo docker-compose up
```

## API

    endpoint: {host:port}/api/v1
* host
```bash
    sudo docker exec -it gingormredisapi_api_1 /bin/bash
    sudo cat /etc/hosts
```
* port

PORT:8000

### API

* `POST`    **/users**    = route for register new user
* `GET`     **/users**    = get data of users
* `DELETE`  **/users/:id**    = delete user with id
* `PUT`     **/users/:id**    = update data with id
* `GET`  **/count**        = get api call counts.


# Swagger Documentation

`host:port/swagger/index.html`

## Install Swagger Library

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

## Generate Swagger Documentation

```bash
sh build.sh
```