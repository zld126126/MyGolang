package util

import mapset "github.com/deckarep/golang-set"

//合并去重数组
func MergeDuplicateIntArray(slice []int, elems []int) []int {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []int
	for i := range t.Iterator().C {
		result = append(result, i.(int))
	}
	return result
}

//数组取出不同元素 放入结果 sourceList中的元素不在sourceList2中 则取到result中
func GetDifferentIntArray(sourceList, sourceList2 []int) (result []int) {
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

//int 数组转换成interface
func IntArrayToInterface(src []int) []interface{} {
	result := []interface{}{}
	for _, v := range src {
		result = append(result, v)
	}
	return result
}

//int 数组去重
func DuplicateIntArray(data []int) []int {
	m := map[int]string{}

	for _, d := range data {
		m[d] = "0"
	}

	l := []int{}
	for key := range m {
		l = append(l, key)
	}
	return l
}

//int array 存在
func ExistInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
