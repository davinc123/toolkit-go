package slice

// ToMap 构造map
func ToMap[T comparable](src []T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))
	for _, val := range src {
		res[val] = struct{}{}
	}

	return res
}

// Deduplicate 确保每个元素不重复
func Deduplicate[T comparable](data []T) []T {
	dataMap := ToMap(data)
	newData := make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
