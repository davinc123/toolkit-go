package slice

// IntersectSet 返回交集
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := ToMap(src)
	res := make([]T, 0, len(src))

	for _, val := range dst {
		if _, ok := srcMap[val]; ok {
			res = append(res, val)
		}
	}
	return Deduplicate(res)
}
