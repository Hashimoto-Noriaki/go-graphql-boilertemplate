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


### データベースのテーブルが作成されているか確認
```
 % psql -h localhost -U user -d mygraphql
```
ログイン後、データベース内のテーブル一覧を表示
```
\dt
```

SQLBoilerを実行
```
sqlboiler postgres
```

### sqlboiler
```
$ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go: downloading github.com/lib/pq v1.10.6
$ sqlboiler --version
SQLBoiler v4.16.2
```
