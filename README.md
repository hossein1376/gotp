# GOTP

This is a simple implementation of an OTP based login system. It was written by
H.Yazdani on August 18th and 19th, 2025.

## Run

To start the application, run:

```shell
docker compose up -d
```

It will listen the `127.0.0.1:9009` by default. To stop it, run:

```shell
docker compose down
```

## Stack

- Go 1.24
- Redis Stack 7.2.0

### Why Redis?

Redis is fast, efficient, simple to use and have great integration in Go. Using
different data structures and features, it can prove itself useful and versatile.

Redis requires no migration or schema definition, which makes it much easier to
experiment with. The schema that RedisSearch uses, also called Index, can be
defined as code and is created at startup, if not exist.

Key expiration allows to the application to store ephemeral keys, such as OTP.
Its fast lookups and auto cleanup are ideal for such scenarios.

Lua script integration allows for atomic operations, combining multiple operations
into one single unit. Effectively, remove any chance of race conditions.

Cheap get and set operations make redis an excellent choice for rate-limit
purposes. This project uses a token-based algorithm to control the number of OTP
requests.

## Open API

The spec file is located [here](assets/openapi.yaml).