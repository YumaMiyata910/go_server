# go_server

This http server is created by golang for testing.

## Start Server

Starting this server with docker.

```bash
docker-compose up --build -d
```

## Call API

There are several endpoints on this server.

```
GET /
GET /v1/greet
POST /v1/post params: {"key": "{value}"}
```
