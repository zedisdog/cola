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