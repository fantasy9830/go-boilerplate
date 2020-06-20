# go-boilerplate

## Configuration

```bash
cp .env.example .env

cp config.yml.example config.yml

cp docker/mysql/docker-entrypoint-initdb.d/init.sql.example docker/mysql/docker-entrypoint-initdb.d/init.sql
```

### Run container

```bash
docker-compose up -d
```

### Enter the app container

```bash
docker-compose exec app bash
```

## App container

### Start

```bash
go run main.go start
```

### Build

```bash
make build
```

### Clean

```bash
make clean
```
