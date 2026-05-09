# templer
go製cliテンプレートエンジン「テンプラ🍤」

# インストール

### Supported platforms
各プラットフォームの最新リリース
|OS |archtechure |
|-|-|
|windows |[amd64](https://github.com/hisuie08/templer/releases/latest/download/templer_windows_amd64.exe)|
|linux |[amd64](https://github.com/hisuie08/templer/releases/latest/download/templer_linux_amd64)|
|linux |[arm64](https://github.com/hisuie08/templer/releases/latest/download/templer_linux_arm64)|

過去のバージョンは [releases](https://github.com/hisuie08/templer/releases) から入手可能

必要に応じて、ダウンロードしたバイナリに実行権限を付与

`chmod +x <downloaded file>`.

## 使い方
### 基本コマンド形式
```
templer <template> [flags]
```

`<template>` には文字列、ファイルパス、ディレクトリを指定可能


## ビルド
リポジトリをクローン
```
make build
```

### example usage
```
mv ./build/templer /sample
cd sample
./templer ./template.tmpl --data ./data.yml
```
