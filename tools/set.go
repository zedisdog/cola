package tools

type Set struct {
	count int
	data  map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		data: make(map[interface{}]bool),
	}
}

func (s *Set) Put(item interface{}) {
	s.data[item] = true
	s.count++
}

func (s *Set) Remove(item interface{}) {
	s.data[item] = false
	s.count--
}

func (s *Set) ToSlice() (result []interface{}) {
	result = make([]interface{}, 0, s.count)
	for key, value := range s.data {
		if value {
			result = append(result, key)
		}
	}
	return
}
