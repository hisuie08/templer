# templer
go製cliテンプレートエンジン「テンプラ🍤」

# Installation


### Supported platforms

|OS|archtechure|
|-|-|
|windows|[amd64](https://github.com/hisuie08/templer/releases/latest/download/windows_amd64.exe)|
|linux|[amd64](https://github.com/hisuie08/templer/releases/latest/download/linux_amd64)|
|linux|[arm64](https://github.com/hisuie08/templer/releases/latest/download/templer_linux_arm64)|

Or you can find [previous releases](https://github.com/hisuie08/templer/releases)

In either case, you may need to run 

`chmod +x <downloaded file>`.

## build yourself
clone repogitory and run
```
make build
```

# example usage
```
mv ./build/templer /sample
cd sample
./templer ./template.tmpl --data ./data.yml
```
