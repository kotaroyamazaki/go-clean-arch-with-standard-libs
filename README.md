# go-clean-arch-with-standard-libs

Golang の標準パッケージを用いた Clean Archtecture + 軽量 DDD のサンプル

# 要件

```
# 想定機能
- 書籍購入サービスにおけるプロフィール登録機能
- プロフィールの項目は以下
  - ニックネーム
  - 生年月日
  - 興味のある本

## 内容
- プロフィールを登録する API の実装
- プロフィールを保持するための RDB のスキーマ設計
  - 設計範囲は以下
    - テーブル名
    - カラム名
    - データ型
    - primary key
    - index
```

## DB のスキーマ設計

- [schema.sql](./schema/schema.sql) に記載

`users`, `books`, `user_favorite_books` テーブルを設計。

## API

- [/internal](./internal/) 以下に実装。REST とクリーンアーキテクチャを用いて構成しました。

POST /users でニックネーム、生年月日、興味のある本の ID を受け取りユーザーの DB への挿入を行います。
興味のある本はすでに DB に本の情報があると仮定して、ID のみを受け取り、`user_favorite_books`テーブル へ挿入することを想定しています。

参考：https://github.com/kotaroyamazaki/go-clean-arch-sample-with-standardissues/1

### 値オブジェクト、エンティティ

`/internal/entites` 以下にエンティティと値オブジェクトを実装。
エンティティ(entity) と 値オブジェクト(VO)は同一パッケージに入れています。
ディレクトリ名と同じファイルがエンティティ、そうでない場合は VO としています。
Golang の特性上コードによる制約は厳密に行っていません。運用ルールによる運用を想定しています。
（例えば 「user entity のみが持つはずの nickename 型が単体で生成ができてしまう。」、「entites.User{}ができてしまう。」など。）

## 概要

| 項目                 | 説明                                                                                               |
| -------------------- | -------------------------------------------------------------------------------------------------- |
| Golang バージョン    | [1.18](https://golang.org/doc/go1.18)                                                              |
| データベース         | 未定義。RDB のものを想定してインメモリで実装。                                                     |
| API                  | REST                                                                                               |
| アーキテクチャ       | [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) |
| ディレクトリ構成参考 | [Standard Go Project Layout](https://github.com/golang-standards/project-layout),                  |
| テスト               | [gotests](https://github.com/cweill/gotests)のフォーマットに沿ったテーブルドリブンテスト           |

# 動作確認

## API

```
curl -X "POST" "http://localhost:8080/users" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d $'{
        "nickname":"takashi",
        "birth_date":"1990-01-01T00:00:00.000+09:00",
        "favorite_book_ids":[1,3,4]
    }'
```

`{"status": "created"}` が返ってくれば OK

## test

```
go test ./...
```

# アーキテクチャ

[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) を元に設計しました。

`entity層`、`usecase層`、`handler層`、`infrastructure層` にわけました。

```
itnernal/
├── entities
│   ├── book
│   ├── repository
│   └── user
├── handlers
├── infra
│   ├── db
│   └── repository
├── middleware
└── usecases
```

![arch](https://user-images.githubusercontent.com/7589567/179364458-4d3b6f93-608f-4273-b793-e7227b05fe2d.png)
