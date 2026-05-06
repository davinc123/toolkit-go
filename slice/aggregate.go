package slice

import "github.com/davinc123/toolkit-go"

// Max 返回最大值
func Max[T toolkit.RealNumber](src []T) T {
	res := src[0]
	for i := 1; i < len(src)-1; i++ {
		if src[i] > res {
			res = src[i]
		}
	}
	return res
}

// Min 返回最小值
func Min[T toolkit.RealNumber](src []T) T {
	res := src[0]
	for i := 1; i < len(src)-1; i++ {
		if src[i] < res {
			res = src[i]
		}
	}

	return res
}

// Sum 返回和
func Sum[T toolkit.RealNumber](src []T) T {
	var res T
	for _, v := range src {
		res += v
	}
	return res
}
