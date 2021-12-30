package dto

const ListOptsTemp = `type ListOptions struct {
	Page    int               ` + "`" + `json:"page"` + "`" + `
	Size    int               ` + "`" + `json:"size"` + "`" + `
	Filters map[string]string ` + "`" + `json:"filters"` + "`" + `
	Limit   int               ` + "`" + `json:"limit"` + "`" + `
}
`
