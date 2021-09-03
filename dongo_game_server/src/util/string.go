package util

import mapset "github.com/deckarep/golang-set"

//数组取出不同元素 放入结果 sourceList中的元素不在sourceList2中 则取到result中
func GetDifferentStrArray(sourceList, sourceList2 []string) (result []string) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

//合并两个字符串数组并去重
func MergeDuplicateStrArray(slice []string, elems []string) []string {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []string
	for i := range t.Iterator().C {
		result = append(result, i.(string))
	}
	return result
}

//string 数组转换成interface
func StrArrayToInterface(src []string) []interface{} {
	result := []interface{}{}
	for _, v := range src {
		result = append(result, v)
	}
	return result
}

//string 数组去重
func DuplicateStrArray(data []string) []string {
	m := map[string]string{}

	for _, d := range data {
		m[d] = "0"
	}

	l := []string{}
	for key := range m {
		l = append(l, key)
	}
	return l
}

//string array 存在
func ExistStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
