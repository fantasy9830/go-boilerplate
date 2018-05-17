利用gin、gorm和viper建構的樣板

# Installation
```bash
$ go get -u github.com/fantasy9830/go-boilerplate
```

```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

## 套件

* [dep](https://github.com/golang/dep)
```bash
$ dep init
```

## Running Migrations
取消註解
```go
migrations.Run()
```

## Usage Example
```http
http://localhost:8080/ping
```

## grpc
啟動grpc server
```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

```bash
$ go run grpc/main.go
```

開啟瀏覽器
```http
http://localhost:8080/grpc
```
可以看到 `Hello your name` 表示成功

## docker
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

開啟瀏覽器
```http
http://localhost:8080/ping
```
可以看到 `pong` 表示成功