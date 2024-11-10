
### コマンド
```
go mod tidy
```

### docker起動
PostgreSQLを5432ポートでホストに公開。
```
docker compose up -d
```

### サーバー起動
```
go run ./server.go
```
### playground
- http://localhost:8080/

<img width="1437" alt="スクリーンショット 2024-11-10 22 35 03" src="https://github.com/user-attachments/assets/c341b4e7-aa00-4a3e-884a-f00256e617c3">

<img width="1439" alt="スクリーンショット 2024-11-10 22 34 41" src="https://github.com/user-attachments/assets/908f3f85-1a0a-45f0-8c34-e479c5fc6915">


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
