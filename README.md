# Word Bubble Poping

## Description
子供向けの英単語ゲームです．

## Requirement
- Docker Compose

## Usage
1. db/.envを作成し，以下の内容を記述してください．
```
    POSTGRES_USER=${ユーザ名}
    POSTGRES_PASSWORD=${パスワード}
    POSTGRES_DB=${データベース名}
```

2. pgadmin/.envを作成し，以下の内容を記述してください．
```
    PGADMIN_DEFAULT_EMAIL=${メールアドレス}
    PGADMIN_DEFAULT_PASSWORD=${パスワード}
```

3. backend/app/.envを作成し，以下の内容を記述してください．
```
    DB_HOST=db
    DB_USER=${db/.envで設定したユーザ名}
    DB_PASSWORD=${db/.envで設定したパスワード}
    DB_NAME=${db/.envで設定したデータベース名}
    DB_PORT=5432
```

4. プロジェクトルートで以下のコマンドを実行してください．
```
    docker-compose up --build
```