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
