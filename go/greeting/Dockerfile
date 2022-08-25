# ベースイメージを指定する
FROM golang:1.19-alpine3.16

# 作業ディレクトリを指定する
# 存在しない場合は作成される
WORKDIR /app

# ローカルにあるファイルをイメージのファイルシステムにコピーする
COPY main.go /app/

# コンテナ実行時の初期設定を指定する
CMD ["go", "run", "main.go"]
