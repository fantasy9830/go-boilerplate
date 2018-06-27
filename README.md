利用gin、gorm和viper建構的樣板

# Installation
```bash
$ go get -u github.com/fantasy9830/go-boilerplate
```

# Start
```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

```bash
$ go run main.go
```

# Running Migrations
建立table
```http
POST http://localhost:8080/migrate/run
```

# Running Seeds
建立user假資料
```http
POST http://localhost:8080/seed/run
```

# Rollback all database migrations
刪除所有migrations
```http
DELETE http://localhost:8080/migrate/reset
```

# grpc
啟動grpc server
```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

```bash
$ go run grpc/main.go
```

```http
GET http://localhost:8080/grpc
```
可以看到 `Hello your name` 表示成功

# docker
```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

build image
```bash
docker build -t go-boilerplate .
```

run
```bash
docker run --rm -p 8080:8080 go-boilerplate
```

```http
GET http://localhost:8080/ping
```
可以看到 `pong` 表示成功