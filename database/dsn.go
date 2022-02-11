package database

import (
	"github.com/zedisdog/cola/errx"
	"github.com/zedisdog/cola/tools"
	"regexp"
)

type Type string

const (
	TypeMysql    Type = "mysql"
	TypePostgres Type = "postgres"
)

type DSN string

func (d DSN) Encode() string {
	return tools.EncodeQuery(string(d))
}

func (d DSN) split() []string {
	reg := regexp.MustCompile(`(^\S+)://(\S+$)`)
	info := reg.FindStringSubmatch(string(d))
	if len(info) < 3 {
		panic(errx.New("dsn is invalid, forget schema?"))
	}
	return info[1:]
}

func (d DSN) Type() Type {
	return Type(d.split()[0])
}

func (d DSN) RemoveSchema() string {
	return d.split()[1]
}
