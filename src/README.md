# 要件

```
# 想定機能
- EC サービスにおけるプロフィール登録機能
- プロフィールの項目は以下
  - ニックネーム
  - 生年月日
  - 興味のある本
# 成果物提出
- 想定機能をどのような設計・コードで実装するか、確認させてください
## 提出物の内容
- プロフィールを登録する API の実装
- プロフィールを保持するための RDB のスキーマ設計
  - 設計範囲は以下
    - テーブル名
    - カラム名
    - データ型
    - primary key
    - index
## 提出期限
- 2022/07/18 23:59:59
```

## DB のスキーマ設計

- [schema.sql](./schema/schema.sql) に記載

users, products, user_favorite_products テーブルを設計しました。

## API

- [/internal](./internal/) 以下に実装

## 設計思想

- 外部ライブラリを使わないようにしました
- DB は用意せずインメモリで行いました
  - RDB の連携部分やトランザクション、ORM のモデル定義などはモックで実装したり、必要最低限で実装するようにしました
- テストはユースケースと entity など重要な部分を中心に書きました
- 一般的に今後必須になるというような部分は現状必要なくても書くようにしました

## 概要

| 項目                 | 説明                                                                                               |
| -------------------- | -------------------------------------------------------------------------------------------------- |
| Golang バージョン    | [1.18](https://golang.org/doc/go1.18)                                                              |
| データベース         | 未定義。RDB のものを想定してインメモリで実装。                                                     |
| API                  | REST                                                                                               |
| アーキテクチャ       | [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) |
| ディレクトリ構成参考 | [Standard Go Project Layout](https://github.com/golang-standards/project-layout),                  |
| テスト               | [gotests](https://github.com/cweill/gotests)のフォーマットに沿ったテーブルドリブンテスト           |

### 補足に対して

| 内容                                                                          | 回答                                                                                                                                                          |
| ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 依存性逆転の法則をどこかで用いてください                                      | README 下部記載のアーキテクチャ図のように DIP を用いてそれぞれの層は内部の層のインターフェースに依存するように実装しました                                    |
| 値オブジェクト、エンティティをどこかで用いてください                          | /internal/entites 以下にエンティティと値オブジェクトを実装しました。Golang の特性上コードによる制約は厳密に行っていません。ルールによる運用を想定しています。 |
| 以下のような追加対応をしていただいても問題ありません（必須ではありません）(略 | 実装したもの：REST, テスト, 軽量 DDD, クリーンアーキテクチャ など                                                                                             |

# 動作確認

## API

```
curl -X "POST" "http://localhost:8080/users" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d $'{
        "nickname":"takashi",
        "birth_date":"1996-03-19T00:00:00.000+09:00",
        "favorite_product_ids":[1,3,4]
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
│   ├── product
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

# 備考

- `/src`というディレクトリ配下にコードを記載しましたが、[参考構成](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#src)では go-like ではないことからアンチパターンとなっています。例外的にホームの [README](../README.md) を汚したくなかったためこのように分けました。
