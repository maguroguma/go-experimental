
```sh
# DB に接続する
docker exec -it postgres-container psql -U myuser -d mydb
# init sql が動かない場合は、ボリュームを削除する必要がある場合がある
docker-compose down -v

# sqlc をインストールする
brew install sqlc

# go get
go get github.com/lib/pq
```
