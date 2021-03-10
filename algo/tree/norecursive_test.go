package ics4

import (
	"reflect"
	"testing"
)

func Test_maxDepth_new(t *testing.T) {
	n2 := &TreeNode{Left: nil, Right: nil, Val: 2}
	n3 := &TreeNode{Left: nil, Right: nil, Val: 3}
	n4 := &TreeNode{Left: nil, Right: nil, Val: 4}
	n1 := &TreeNode{Left: n3, Right: n4, Val: 1}
	n0 := &TreeNode{Left: n1, Right: n2, Val: 0}
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"测试1", args{n0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth_new(tt.args.node); got != tt.want {
				t.Errorf("maxDepth_new() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthBfs(t *testing.T) {
	n2 := &TreeNode{Left: nil, Right: nil, Val: 2}
	n3 := &TreeNode{Left: nil, Right: nil, Val: 3}
	n4 := &TreeNode{Left: nil, Right: nil, Val: 4}
	n1 := &TreeNode{Left: n3, Right: n4, Val: 1}
	n0 := &TreeNode{Left: n1, Right: n2, Val: 0}

	ans := [][]int{{1}}
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{n0}, ans},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthBfs(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthBfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
