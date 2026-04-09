package slice

import (
	"fmt"
	"testing"
)

func intEqual(a, b int) bool { return a == b }

func TestDeleteIndex(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		index   int
		want    []int
		wantErr bool
	}{
		{name: "删除头部元素", src: []int{1, 2, 3}, index: 0, want: []int{2, 3}},
		{name: "删除中间元素", src: []int{1, 2, 3}, index: 1, want: []int{1, 3}},
		{name: "删除尾部元素", src: []int{1, 2, 3}, index: 2, want: []int{1, 2}},
		{name: "index为负数", src: []int{1, 2, 3}, index: -1, wantErr: true},
		{name: "index越界", src: []int{1, 2, 3}, index: 3, wantErr: true},
		{name: "index远越界", src: []int{1, 2, 3}, index: 99, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteIndex(tt.src, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("DeleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteElement(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		element int
		want    []int
	}{
		{name: "删除存在的单个元素", src: []int{1, 2, 3}, element: 2, want: []int{1, 3}},
		{name: "删除所有重复元素", src: []int{1, 2, 2, 3, 2}, element: 2, want: []int{1, 3}},
		{name: "删除不存在的元素", src: []int{1, 2, 3}, element: 9, want: []int{1, 2, 3}},
		{name: "删除头部元素", src: []int{1, 2, 3}, element: 1, want: []int{2, 3}},
		{name: "删除尾部元素", src: []int{1, 2, 3}, element: 3, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteElement(tt.src, tt.element, intEqual)
			if fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("DeleteElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFirstElement(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		element int
		want    []int
	}{
		{name: "删除第一个匹配元素", src: []int{1, 2, 3, 2}, element: 2, want: []int{1, 3, 2}},
		{name: "只有一个匹配", src: []int{1, 2, 3}, element: 2, want: []int{1, 3}},
		{name: "不存在的元素", src: []int{1, 2, 3}, element: 9, want: []int{1, 2, 3}},
		{name: "删除头部元素", src: []int{1, 2, 3}, element: 1, want: []int{2, 3}},
		{name: "删除尾部元素", src: []int{1, 2, 3}, element: 3, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteFirstElement(tt.src, tt.element, intEqual)
			if fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("DeleteFirstElement() = %v, want %v", got, tt.want)
			}
		})
	}
}