# cola

## requirement
go version > 1.16

## install
```bash
go install github.com/zedisdog/cola/cmd/cola@latest
```
## usage
```bash
cola new test #this will install the package github.com/golang-migrate/migrate/v4/cmd/migrate@latest
cd test
docker-compose up -d #start mysql and redis
cd test/cmd/app
go run main.go
```