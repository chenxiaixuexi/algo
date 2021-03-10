package ics4

import "testing"

func TestFindMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"ceshiyi", args{[]int{2, 1}}, 1},
		{"ceshi2", args{[]int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"ceshi3", args{[]int{11, 13, 15, 17}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMin(tt.args.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
