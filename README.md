# GoでCRUDなWebAppのサンプル（身内用）
##  開発環境
os:os x high sierra

ide:goland

##  フレームワーク＆ライブラリ

gin + gorm

##  手順
* goのインストール

```
brew install go
```

* 確認
```
go version
```

* Goのワークスペースを作成

  自分は /Users/admin/go に作成している

* PATH設定を ~/.zprofile もしくは ~/.bash_profile に追加

  GOPATHとPATHが通っていないと、後にエラーでます

```
vi ~/.bash_profile

export GOPATH=$HOME/go //さっき作ったワークスペースに通す

export PATH=$PATH:$(go env GOPATH)/bin

source ~/.bash_profile
```

 　 go env と printenv PATH  で設定が反映されてるか確認する

## 実行

   ターミナルを開いてローカルに落としたこのリポジトリに  cd

   mysql の server 起動させて

   以下の順にコマンドを打つ

```
make deps

make migrate/init

make migrate/up

make run
```

### 以上、ザックリ書いたから動かなければ聞きに来て



