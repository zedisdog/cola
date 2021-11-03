# cola

## requirement
go version >= 1.17

## install
```bash
go install github.com/zedisdog/cola/cmd/cola@develop
```
## usage
```bash
cola new test #suggest install github.com/golang-migrate/migrate/v4/cmd/migrate@latest for database migrate
cd test
docker-compose up -d #start mysql and redis and app
```
### settings in environment variable
cause shell not supports the environment variable which contains `.`, use `__` instead.
e.g.: `WX__WECHAT__APPID` is presents `wx.wechat.appId` or `wx.wechat.appid`. ps: keys is case-insensitive in viper.

### seed
> note: all seed function will be wrapped by transaction

### swagger
builtin [`swagger-ui` 3.52.5](https://github.com/swagger-api/swagger-ui)

gin for example:
```go
r := gin.Default()
//set a router to serve the spec file.
//assume the url of this router is http://localhost/swagger.
r.StaticFile("swagger", "path/to/swagger.yaml")
// set a router for swagger-ui
r.StaticFS(
    "swaggerui",
    // use swagger.SwaggerUI() can get a fs.FS instance, which set the spec file uri to specific one.
    http.FS(swagger.SwaggerUI("http://localhost/swagger")),
)
```