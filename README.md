# Docker-webapp

## このレポジトリについて
GolangおよびDockerを用いた簡単なREST APIを作成しました．
Userに関するName, Emailのカラムの登録および取得が可能になっています．  

## ディレクトリ構成


```
root/
　├ postgres/  Postgresqlを初期化するためのsqlのスクリプト
　│
　├ .env   localで実行する際のポート番号などの環境変数
　│
　├ Dockerfile   docker立ち上げに必要な処理など
　│
　├ docker-compose.yml   dbとappを同時に立ち上げる
　│
　└ main.go  REST APIをここで提供　　　
```

## ローカル環境でのビルド
手元の環境でビルドを行うにはrootディレクトリで以下のコマンドを実行します．
```
$ docker-compose build
$ docker-compose up -d
```
