package tpl

import (
	"bytes"
	"text/template"
)

const UpTableTemp = `CREATE TABLE {{.}} (
    id BIGINT(20) UNSIGNED PRIMARY KEY,
    created_at BIGINT(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    updated_at BIGINT(20) NOT NULL DEFAULT 0 COMMENT '更新时间'
)ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='' COLLATE=utf8mb4_unicode_ci;
`

const DownTableTemp = "DROP table `{{.}}`;"

var ReservedKeyword = []string{}

func GenMigration(tableName string) (up []byte, down []byte, err error) {
	tmp, err := template.New("up").Parse(UpTableTemp)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, tableName)
	if err != nil {
		return
	}
	up = buff.Bytes()

	tmp, err = template.New("down").Parse(DownTableTemp)
	if err != nil {
		return
	}
	buff = bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, tableName)
	if err != nil {
		return
	}
	down = buff.Bytes()

	return
}
