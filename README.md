# dab
domestic accounts book app

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

## TODOs
### Backend
- [ ] testの追加
- [ ] RequestBodyのプロパティ欠損時の対応
- [ ] Middlewareの調査
- [ ] Loggerの追加

### Frontend
- [x] Reactの導入
- [x] React Routerの導入
- [ ] React Hook Formの導入？
- [x] MUIの導入
- [x] prettierの設定
- [ ] ディレクトリ構成の決定
- [ ] Storybookの導入？
- [ ] Recoilの導入?

## Notes
- CHFだと浮動小数点数で金額を扱う必要が出てきそう
  - 厳密な金額管理が必要なわけではないので無視
- 税金どうするか
  - きちんと管理したいわけではないので税込みで入力する
- REST APIの設計においてGET /receiptsで複数の検索条件をクエリパラメータにするのが気になった
  - [Design RESTful query API with a long list of query parameters](https://stackoverflow.com/questions/14202257/design-restful-query-api-with-a-long-list-of-query-parameters)を参考にして、POST /receipts/searchとした
- backend/testing/data/db/test.sqlから日本語のカテゴリー名を流し込むと文字化けしてしまった
  - 直接mysqlを除くと日本語で表示されていたのでGo側の文字コード等に問題があるのかと思ったが、実際はsqlを記述する際にn-prefixというものが必要なだけであった
  - 参考：[SQLServerに日本語データをInsertすると文字化けする。](https://kitigai.hatenablog.com/entry/2018/05/27/010440)

## Libraries
### Backend
- [Echo](https://echo.labstack.com/)
- [migrate](https://github.com/golang-migrate/migrate)
- [xo](https://github.com/xo/xo)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [A5:SQL Mk-2](https://a5m2.mmatsubara.com/)

### Frontend
- [Vite](https://ja.vitejs.dev/)
- [React](https://react.dev/)
- [React Router](https://reactrouter.com/en/main)
- [React Hook Form](https://react-hook-form.com/)
- [MUI](https://mui.com/)
- [AXIOS](https://axios-http.com/)
- [Orval](https://orval.dev/)

