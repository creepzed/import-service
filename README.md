# import-services-mirakl

> Service to import data from a csv from Mirakl.


# Requirements

- [Docker-compose](https://docs.docker.com/compose/)
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/)

## Usage on Development
1. Copy .env.example to .env
2. Execute command.
    ```sh
    go run app/main.go
    ```
3. Happy Coding.

## Usage with docker-compose
1. Copy .env.example to .env
2. Execute command.
    ```sh
    docker-compose up -d
    ```
3. Happy Coding.

## Default port services
- **Zookeeper:** 9000
- **Kafka UI:** 9001
- **Database (Postgres):** 9002
- **Admin Database (Adminer):** 9003
- **Kafka:** 9004