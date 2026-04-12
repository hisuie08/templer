# templer
go製cliテンプレートエンジン「テンプラ🍤」

# installation

```
go build -o ./build/templer .
chmod +x ./build/templer
```

# example usage
```
mv ./build/templer /sample
cd sample
./templer ./template.tmpl --data ./data.yml
```
