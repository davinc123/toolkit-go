package slice

import "github.com/davinc123/toolkit-go/internal/slice"

// Delete 删除 index 处的元素
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, _, err := slice.Delete[Src](src, index)
	return res, err
}

// FilterDelete 删除符合条件的元素
func FilterDelete[T any](src []T, filter func(index int) bool) []T {
	emptyPos := 0
	for idx := range src {
		if filter(idx) {
			continue
		}
		src[emptyPos] = src[idx]
		emptyPos++
	}

	return src[:emptyPos]
}
