# ベースイメージを指定する
FROM python:3.9-slim

# 作業ディレクトリを指定する
# 存在しない場合は作成される
WORKDIR /app

# ローカルにあるファイルをイメージのファイルシステムにコピーする
COPY main.py /app/

# コンテナ実行時の初期設定を指定する
CMD ["python", "main.py"]
