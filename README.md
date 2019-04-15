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
その後，以下のURLにアクセスします．
http://localhost:8080/

## 想定しているリクエストおよびその結果
### Hello, worldの表示
リクエスト  
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
レスポンス(json)
```
{"message":"Hello, World!!"}
```

### DBに登録されているユーザ一覧を取得
リクエスト  
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```
レスポンス(json)
```
[{"id":1,"name":"sugiyama","email":"fuga.com"},{"id":2,"name":"tanaka","email":"hoge.com"}]
```

### DBに登録されているユーザ情報をIDで指定
リクエスト  
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```
レスポンス(json)
```
{"id":1,"name":"sugiyama","email":"fuga.com"}]
```

### DBに新しいユーザを登録
リクエスト  
```
$ curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```
レスポンス(json)
```
{"id":3,"name":"test","email":"hoge@example.com"}
```
このあと/usersにアクセスすると追加されていることがわかる．
```
[{"id":1,"name":"sugiyama","email":"fuga.com"},{"id":2,"name":"tanaka","email":"hoge.com"},{"id":3,"name":"test","email":"hoge@example.com"}]
```
(値の更新および削除は時間がなかったため省略しました)  

## Herokuへのデプロイ
今回はDockerを簡単にデプロイすることのできるPaaSとしてHerokuを選択した．  
以下にデプロイまでの手順を示す． 
