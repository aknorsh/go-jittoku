# go-jittoku

> GoLand にて動作検証
>
> go.mod をこのリポジトリ内に配置する場合、go get しないと module が解決できないことが分かった。のでお蔵入りかな〜

## Usage

たとえば、こういう状況

```text
go_prj_root
- go.mod
- go.sum
- cmd
- src
  - hoge_service.go <- ここで開発中
  - fuga_service.go
- util
```

- STEP1: デバッグ用ディレクトリをつくり、中身を全部 ignore

```sh
cd go_prj_root

mkdir .debug
cd .debug
touch .gitignore
echo "*" > .gitignore
```

- STEP2: デバッグ用ディレクトリ以下に git clone

以上。これで go.mod を汚さずにリポジトリの関数が使えるはず。

```text
go_prj_root
- .debug
  - .gitignore
  - go-jittoku
    - README.md
    - go.mod
    - print.go <- ここの関数を使いたければ、
    ...
- go.mod
- go.sum
- cmd
- src
  - hoge_service.go <- ここでimportするだけでヨシ
  - fuga_service.go
- util
```
