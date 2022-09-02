# swaggofix
将swagger.json中的/:path变更为/{:path}

## 场景

beego在生成路由的时候，代码中对路由的定义需要是 `/:path` 形式，但[swaggo](https://github.com/swaggo/swag)按这个规则生成的swagger.json无法用于调试，因此使用此工具将路由调整为`/{:path}`的形式


## 使用方法

```
go install github.com/fanhousanbu/swaggofix@latest

swaggofix --path=[swagger.json]路程
```
