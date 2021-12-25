package mysql

type CommonField struct {
	ID        uint64 `json:"id,string"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
