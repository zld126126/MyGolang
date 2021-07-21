package util

//interface 去重
func DuplicateInterface(data []interface{}) []interface{} {
	m := map[interface{}]interface{}{}

	for _, d := range data {
		m[d] = "0"
	}

	l := []interface{}{}
	for key := range m {
		l = append(l, key)
	}
	return l
}

//interface 存在
func ExistInterface(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
