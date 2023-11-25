# misskey クライアントを Go で作ってみるテスト

## Introduction

[misskey](https://join.misskey.page/ja-JP/)のエンドポイントを Go で実行してみようと出来心で思い立って実装してみました。

## Features

### できたこと

現在は CUI 含めてユーザ入力を受け付けるインターフェースがないため、すべてテストコードから実行する必要があります。  
Dev Container を採用しているため、実行に際して Go のインストールは必須ではありません。(VSCode, Docker は必要です)

- `/meta` エンドポイントの実行
- `/notes/create` エンドポイントの実行とノートの作成
- API キー、ベース URL は設定ファイル(yaml)に退避

### 実行手順

WIP

## TODO

### やること

- cobra の導入と cli 化
  - アカウントの追加(yaml->toml)
  - 実装済みエンドポイントの cli 対応
- ホームタイムラインの取得
  - tl エンドポイント
  - streaming api 対応
- 通知(notifications)の取得

### 今後の予定

- リアクションの取得
- カスタム絵文字はいったん諦めましょう

## 参考

https://misskey.io/cli

https://github.com/mikuta0407/misskey-cli/tree/main
