# RabbitMQ connection checker in Docker

## Usages

```sh
# localhost:5672, no auth
docker run --rm --network host ttulka/rabbitmq-connection-checker
```

```sh
# localhost:5672, with credentials
docker run --rm --network host ttulka/rabbitmq-connection-checker -u user -p pass
```

```sh
# with host and port, no auth
docker run --rm --network host ttulka/rabbitmq-connection-checker --host 10.0.0.1 --port 15672
```

```sh
# with TLS enabled, no auth
docker run --rm --network host ttulka/rabbitmq-connection-checker --tls
```