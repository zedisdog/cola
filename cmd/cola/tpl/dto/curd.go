package dto

const CurdTemp = `type {{.Method}}{{.Entity}} struct {
{{range .Fields}}
	{{.FieldName}}	{{.FieldType}}	` + "`" + `json:"{{.FieldNameSnake}}" binding:"required"` + "`" + `
{{end}}
}
`
