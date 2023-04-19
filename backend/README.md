# dab-backend

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

## References
- [echo.Context を最大限に活用する](https://codehex.hateblo.jp/entry/echo-context)
- [OpenAPIからGo言語のコードを自動生成してくれるツールを試してみた](https://zenn.dev/rescuenow/articles/3c9a19eb2c0655)
- [Go (Echo) で複雑な構造のJSON(POSTリクエスト)からデータを取得したい](https://teratail.com/questions/vfwi04zdo1pkvr)
- [oapi-codegenを使ってみた](https://speakerdeck.com/akeno/oapi-codegenwoshi-tutemita)
- [[Go] OpenAPI コード自動生成でビジネスロジックに集中する開発へ](https://qiita.com/nyanchu/items/1c259750352b49e96a18)
- [golangのxoを導入を決めてファイルの運用方法がいい感じになってきたので書いておく](https://tsuyoshi-nakamura.hatenablog.com/entry/2018/11/16/100133)
- [GoのWebアプリをテストするノウハウ](https://zenn.dev/media_engine/articles/testing-go-applications)
