# go-jittoku

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
