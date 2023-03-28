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
