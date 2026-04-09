package slice

import "errors"

// DeleteIndex 删除 index 的元素
func DeleteIndex[Src any](src []Src, index int) ([]Src, error) {
	if index < 0 || index >= len(src) {
		return src, errors.New("index out of range")
	}

	return append(src[:index], src[index+1:]...), nil
}

// DeleteElement 删除 所有element 元素
func DeleteElement[Src any](src []Src, element Src, equal func(a, b Src) bool) []Src {
	newSrc := make([]Src, 0, len(src)-1)
	for _, v := range src {
		if !equal(v, element) {
			newSrc = append(newSrc, v)
		}
	}
	return newSrc
}

// DeleteFirstElement 删除 第一个 element 元素
func DeleteFirstElement[Src any](src []Src, element Src, equal func(a, b Src) bool) []Src {
	newSrc := src[0:]
	for index, v := range src {
		if equal(v, element) {
			newSrc = append(src[:index], src[index+1:]...)
			break
		}
	}

	return newSrc
}
