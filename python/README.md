# Python の実行環境を作成する

今日の日付を返すスクリプトをコンテナで実行する

## docker コマンド

イメージの作成

```
docker build . -t python3.9
```

コンテナの実行

```
docker run --rm python3.9
```

コンテナの中に入る

```
docker run --rm -it python3.9 /bin/bash
```
