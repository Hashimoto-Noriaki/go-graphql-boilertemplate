### サーバー起動
```
go run ./server.go
```


### コマンド
```
go mod tidy
```

### docker起動
PostgreSQLを5432ポートでホストに公開。
```
docker compose up -d
```

### playground

### sqlboiler
```
$ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go: downloading github.com/lib/pq v1.10.6
$ sqlboiler --version
SQLBoiler v4.16.2
```
