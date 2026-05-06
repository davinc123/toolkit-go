package slice

import "github.com/davinc123/toolkit-go/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}

	res := src[index]

	for i := index; i < length; i++ {
		src[i] = src[i+1]
	}
	return src, res, nil
}
