# Contributing

## Software

- Node.js 20
- Pnpm 8
- Golang 1.22

## Build

```bash
pnpm run build
```

## mamayoについて

mamayoはGolangで記述された管理ツールです。現在、以下のサブコマンドが実装されています。

- make READMEを生成します
- url URLをチェックします
- testcase
  - check 全てのチェックを実行します
  - dup   重複がないかチェックします
  - make  `gomamayo.csv` `gomamayo.txt`を生成します
  - sort  `gomamayo.json`をソートします

## 主に編集するべき場所

### /gomamayo.toml

READMEに掲載されているリポジトリやウェブサイトのデータです。これらのデータを基にmamayoがREADME.mdを生成します。

READMEの生成は`pnpm run build`を実行します。

### /assets/template.md

mamayoがREADME.mdを生成する際のテンプレートです。生成には`text/template`を使用しています。

### /test/gomamayo.json

ゴママヨテストケースです。詳細は[/test/README.md](/test/README.md)を参照してください。
