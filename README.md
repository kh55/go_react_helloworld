## build
docker-compose build

# react-appというプロジェクト名を使用する場合、react-appというフォルダが作成される
docker-compose run --rm react sh -c 'npx create-react-app react-app --template typescript'

# コンテナ直下にreactプロジェクトを作成する場合
docker-compose run --rm react-app sh -c 'npx create-react-app . --template typescript'

# nextjsをインストール
docker-compose run --rm react-app sh -c 'npx create-next-app .'

# env

REDIS_ADDR: redis
REDIS_PORT: 6379
REDIS_DB: 0
REDIST_PASSWORD: ""
