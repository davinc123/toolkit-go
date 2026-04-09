package slice

import "errors"

// Add 在index处添加元素
// index 范围应为[0, len(src)]
// 如果index == len(src) 则表示往末尾添加元素
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	if index < 0 || index > len(src) {
		return src, errors.New("index out of range")
	}

	newSrc := make([]Src, 0, len(src)+1)
	newSrc = append(newSrc, src[:index]...)
	newSrc = append(newSrc, element)
	newSrc = append(newSrc, src[index:]...)

	return newSrc, nil
}
