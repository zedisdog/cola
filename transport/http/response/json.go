package response

type Json string

func (j Json) MarshalJSON() (data []byte, err error) {
	if j == "" {
		return []byte("null"), nil
	}
	return []byte(j), nil
}
