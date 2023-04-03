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
- [x] MUIの導入
- [ ] ディレクトリ構成の決定
- [ ] Storybookの導入？
- [ ] Scaffdogの導入？
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

## References
### Official
#### Backend
- [Echo](https://echo.labstack.com/)
- [migrate](https://github.com/golang-migrate/migrate)
- [xo](https://github.com/xo/xo)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [A5:SQL Mk-2](https://a5m2.mmatsubara.com/)

##### Frontend
- [Vite](https://ja.vitejs.dev/)
- [React](https://react.dev/)
- [React Router](https://reactrouter.com/en/main)
- [MUI](https://mui.com/)
#### Others
- [echo.Context を最大限に活用する](https://codehex.hateblo.jp/entry/echo-context)
- [OpenAPIからGo言語のコードを自動生成してくれるツールを試してみた](https://zenn.dev/rescuenow/articles/3c9a19eb2c0655)
- [Go (Echo) で複雑な構造のJSON(POSTリクエスト)からデータを取得したい](https://teratail.com/questions/vfwi04zdo1pkvr)
- [oapi-codegenを使ってみた](https://speakerdeck.com/akeno/oapi-codegenwoshi-tutemita)
- [[Go] OpenAPI コード自動生成でビジネスロジックに集中する開発へ](https://qiita.com/nyanchu/items/1c259750352b49e96a18)
- [golangのxoを導入を決めてファイルの運用方法がいい感じになってきたので書いておく](https://tsuyoshi-nakamura.hatenablog.com/entry/2018/11/16/100133)
- [GoのWebアプリをテストするノウハウ](https://zenn.dev/media_engine/articles/testing-go-applications)
- [フロントエンドエンジニアが参考にするデザインシステムをまとめてみた](https://qiita.com/bikky_no_yakata/items/368b02c3df178d4e0472)
- [フロントエンド API通信戦略](https://zenn.dev/sutamac/articles/27246dfe1b5a8e)