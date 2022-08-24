# Docker Handson

Dockerのハンズオン資料

## docker コマンド

イメージの作成

```
docker build パス -t イメージ名:タグ名
```

イメージの確認

```
docker images
```

コンテナの実行

```
docker run イメージ名
```

コンテナの一覧を表示する

```
docker ps -a
```

コンテナを停止する

```
docker stop [コンテナID or コンテナ名]
```

コンテナを削除する

```
docker rm [コンテナID or コンテナ名]
```
