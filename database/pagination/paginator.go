package pagination

type Paginator interface {
	Page(container interface{}, currentPage int, pageSize int) (total int, err error)
}
