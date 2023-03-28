# dab
domestic accounts book app

## Structure
```shell
.
├── cmd
│   └── dab
├── ddl
├── docker
│   └── dev
├── gen
│   ├── daocore
│   ├── dbschema
│   ├── ddl
│   └── oapi
├── script
│   └── dbgen
│       └── templates
└── src
    ├── app
    │   ├── api
    │   ├── server
    │   └── service
    ├── dberror
    ├── env
    ├── errors
    └── util
        ├── filter
        └── name
```
- `cmd`
  - サーバーを立てるためのプログラムを配置
- `ddl`
  - RDBのためのDDLを配置
  - go-migrateで使えるような形式
- `docker`
  - `dev`
    - 開発環境用のDockerfile
- `gen`
  - 自動生成されるファイルを配置
  - `daocore`
    - xoで生成したスキーマ情報を用いて生成するDAO
  - `dbschema`
    - xoで生成されるデータベースのスキーマ情報
  - `ddl`
    - [A5:SQL Mk-2](https://a5m2.mmatsubara.com/)のERDから生成したDDL（これを基にgo-migrate用のDDLを記述）
  - `oapi`
    - openAPI Schemaから生成したEcho用の型・インターフェース定義
- `script`
  - `dbgen`
    - xoのschemaからDAOを生成するためのプログラム群
- `src`
  - ソースコード

## Functions to be implemented
### Frontend
- レシートの入力
  - 購入日
  - 購入場所
  - 購入品目
    - 商品名
    - 金額
    - 用途（食費・娯楽費等）
  - 割引額
- 入力されたレシートの一覧表示
  - 日付・金額・場所等で絞り込めるように
- 支出の可視化
  - グラフ

### Backend
- 複数通貨のサポート
  - とりあえずJPY・CHF

## Notes
- CHFだと浮動小数点数で金額を扱う必要が出てきそう
  - 厳密な金額管理が必要なわけではないので無視
- 税金どうするか
  - きちんと管理したいわけではないので税込みで入力する
- REST APIの設計においてGET /receiptsで複数の検索条件をクエリパラメータにするのが気になった
  - [Design RESTful query API with a long list of query parameters](https://stackoverflow.com/questions/14202257/design-restful-query-api-with-a-long-list-of-query-parameters)を参考にして、POST /receipts/searchとした

## References
- [Echo](https://echo.labstack.com/)
- [echo.Context を最大限に活用する](https://codehex.hateblo.jp/entry/echo-context)
- [OpenAPIからGo言語のコードを自動生成してくれるツールを試してみた](https://zenn.dev/rescuenow/articles/3c9a19eb2c0655)
- [Go (Echo) で複雑な構造のJSON(POSTリクエスト)からデータを取得したい](https://teratail.com/questions/vfwi04zdo1pkvr)
- [oapi-codegenを使ってみた](https://speakerdeck.com/akeno/oapi-codegenwoshi-tutemita)
