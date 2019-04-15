# Docker-webapp

## このレポジトリについて
GolangおよびDockerを用いた簡単なREST APIを作成した．
Userに関するName, Emailのカラムの登録および取得が可能になっている．  

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
手元の環境でビルドを行うにはrootディレクトリで以下のコマンドを実行する．
```
$ docker-compose build
$ docker-compose up -d
```
その後，以下のURLにアクセスすることでレスポンスが表示される．  
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
以下にデプロイまでの手順を示す． heroku cliを予めインストールしておく．

1. herokuへログイン
```
$ heroku login
```
2. 新しいappを作成
```
$ heroku create docker-webapp-2
```
3. ローカルのディレクトリをpushし，その結果をheroku側に反映させる
```
$ heroku container:push web --app docker-webapp-2
$ heroku container:release web --app docker-webapp-2
```
4. 公開が完了したのでURLにアクセスする．
今回は以下のURLにサンプルプログラムをデプロイした．  
http://docker-webapp-2.herokuapp.com/

## 想定しているリクエストおよびその結果
ローカルで実行したリクエストのURL部分を "http://docker-webapp-2.herokuapp.com/" に置き換えることで同様の結果を得ることができる．

##　達成したこと
- **ローカルでのdocker-composeを用いたappとdbの同時起動**
- **Postgresqlに対する読み書きの実装**
- **Golangを用いたGET, POSTリクエストの受付およびレスポンスの返却**
- **Herokuへのdocker環境のデプロイ**

## 達成できなかったこと
- **エラー処理**  
あらゆるパターンに対するエラー処理を実装する余裕がなかった． 存在しないIDのデータにアクセスしようとするなどの例外に対応しきれていない．
- **Heroku環境へのPostgresqlの反映**  
docker-composeでPostgresqlのDBをHeroku上に生成するように設定しようとしたが，様々な方法を用いてもうまくいかなかった．  
Heroku上に手動でPostgresqlを作成しそのDBに関する情報を環境変数として持つことで仮の接続状態を作ったが，今回の目標であるdockerによる自動生成はできていない．



