package ics4

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, -1}, 1}, []int{1, -1}},
		{"ceshi2", args{[]int{4, -2}, 2}, []int{4}},
		{"ceshi", args{[]int{4, -2, 5, 6}, 2}, []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
