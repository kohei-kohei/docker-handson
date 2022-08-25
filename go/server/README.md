## docker コマンド

イメージの作成

```
docker build . -t handson-server
```

コンテナの実行

```
docker run --rm -p 8000:8000 handson-server
```

コンテナの中に入る

```
docker run --rm -it handson-server /bin/sh
```
