package slice

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		element int
		index   int
		want    []int
		wantErr bool
	}{
		{name: "在头部插入", src: []int{1, 2, 3}, element: 0, index: 0, want: []int{0, 1, 2, 3}},
		{name: "在中间插入", src: []int{1, 2, 3}, element: 9, index: 1, want: []int{1, 9, 2, 3}},
		{name: "在末尾插入", src: []int{1, 2, 3}, element: 4, index: 3, want: []int{1, 2, 3, 4}},
		{name: "index为负数", src: []int{1, 2, 3}, element: 0, index: -1, wantErr: true},
		{name: "index越界", src: []int{1, 2, 3}, element: 0, index: 99, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.src, tt.element, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
