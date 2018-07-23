# GO 後臺管理系統

前端頁面可使用 [react-boilerplate](https://github.com/fantasy9830/react-boilerplate) 搭配

## Installation
```bash
$ go get -u github.com/fantasy9830/go-boilerplate
```

## Start

需安裝 `PostgreSQL`

設定資料庫帳號密碼 `$GOPATH/src/github.com/fantasy9830/go-boilerplate/config/debug.yaml`

```bash
$ cd $GOPATH/src/github.com/fantasy9830/go-boilerplate
```

```bash
$ go run main.go
```

## Running Migrations
建立table
```http
POST http://localhost:8080/migrate/run
```

## Running Seeds
建立user假資料
```http
POST http://localhost:8080/seed/run
```

## Rollback all database migrations
刪除所有migrations
```http
DELETE http://localhost:8080/migrate/reset
```
## Auth
`需要執行 Migration + Seed`

### Role-based access control
| permissions | 說明                                            |
|-------------|-------------------------------------------------|
| action      | 表示權限，例：get、post、read、write、delete... |
| guard_name  | 表示使用的系統，例：web、api、erp...            |

| roles      | 說明                                 |
|------------|--------------------------------------|
| guard_name | 表示使用的系統，例：web、api、erp... |

## grpc
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

```http
GET http://localhost:8080/ping
```
可以看到 `pong` 表示成功

## Features

* [x] gin
* [x] gorm + migration + seed
* [x] config(viper)
* [x] grpc
* [x] docker
* [x] CORS
* [x] 登入認證功能(JWT)
* [x] 權限管理
* [x] Repository and Services Pattern
* [ ] ...
