package slice

import "github.com/davinc123/toolkit-go/internal/slice"

func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add(src, element, index)
	return res, err
}
