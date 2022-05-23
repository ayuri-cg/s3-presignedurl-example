# s3-presignedurl-example

aws s3 presigned url(PUT)のsample code

# 事前準備

1. direnvをインストールしておく
2. main.goのバケット名、ファイル名は適宜修正しておく
3. 実行するアカウントには対象bucketに対し、PutObjectの権限を付与しておく

# How to use

## 初回のみ
1. .envrc.tmpl --> .envrcとrename
2. 環境変数を入力
3. direnv allowを実行して反映しておく

## 実行方法
1. go run main.goを実行
2. 出力されたURLを使い、RESTツールなどで、PUTリクエストを実行する
   - METHOD: PUT
   - HEADER: Content-Type: application/pdf
   -ー> ファイル形式に合わせる 
   - URL: 出力されたURLをセット
   - BODY: アップロード対象ファイルをセット
3. S3バケットにファイルがアップロードされていることを確認する