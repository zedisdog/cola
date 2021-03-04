package response

type Json string

func (j Json) MarshalJSON() (data []byte, err error) {
	return []byte(j), nil
}
