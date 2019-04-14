FROM golang:1.9

# コンテナ作業ディレクトリの変更
WORKDIR /go/src/denki/go
# ホストOSの ./go の中身を作業ディレクトリに追加
ADD ./go .
