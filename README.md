# popular

サイト URL:https://popular-32pe64nwja-an.a.run.app

## 概要

ツイートができて、そこにフリー画像の写真を検索して投稿できるアプリ

## 初期設定

<p>Go ModuleやMakefileなど、詳細はSETTING.mdに記載</p>

## 開発環境
docker version 20.10.6が必要。なければ下記URLからダウンロード。  
https://hub.docker.com/

認証機能を利用するためにはTwitterDeveloperでプロジェクトを作成し、ルートディレクトリのdocker/以下に.envを作成、環境変数を設定する。  
プロジェクト作成の参考記事  
https://qiita.com/wheatandcat/items/fe66c7ee2521a6966505  

```.env  
TWITTER_KEY=xxxxxxxx
TWITTER_SECRET=xxxxxxxx
```

写真検索にはPixabay APIを利用するため下記記事参照。.envに環境変数追加   
https://www.whizz-tech.co.jp/5449/

```.env
PIXABAY_KEY=xxxxxxxx
```

1. gitリポジトリをクローンする
```
git clone https://github.com/keigooba/popular.git
```
2. リポジトリのルートディレクトリからdocker起動(port:8080が必要)
```
cd docker && docker-compose up -d
```
3. dockerコンテナに入る
```
docker-compose exec popular bash
```

## 機能

## 技術

フロントエンド・バックエンド
1. Go1.16.2
2. Bootstrap4.5.0
3. jQuery3.5.1

インフラ
1. cloud Build + cloud Run
2. Datastore(NoSQL)予定
