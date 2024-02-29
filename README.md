# コンテナ
go
react
redis
phpredisadmin

## build
docker-compose build

## react-appというプロジェクト名を使用する場合、react-appというフォルダが作成される
docker-compose run --rm react sh -c 'npx create-react-app react-app --template typescript'

## コンテナ直下にreactプロジェクトを作成する場合
docker-compose run --rm react-app sh -c 'npx create-react-app . --template typescript'

## nextjsをインストール
docker-compose run --rm react-app sh -c 'npx create-next-app .'

# go/.env
```
REDIS_ADDR: redis
REDIS_PORT: 6379
REDIS_DB: 0
REDIST_PASSWORD: ""
```

# Nextjs
### / 
Next.js初期画面

### /api-test
go-api連携テスト

### input-test
入力した文字列をgo-apiに送信してredisに登録
返却される入力したデータを受け取り表示する

# go API
### / 
index

### /hello
hello worldを返す

### redis
redisにデータを登録

### input
受信したデータをredisに登録し、データを返す
