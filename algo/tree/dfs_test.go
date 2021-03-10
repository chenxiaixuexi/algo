package ics4

import "testing"

func Test_maxDepth(t *testing.T) {
	n2 := &TreeNode{Left: nil, Right: nil, Val: 2}
	n3 := &TreeNode{Left: nil, Right: nil, Val: 3}
	n4 := &TreeNode{Left: nil, Right: nil, Val: 4}
	n1 := &TreeNode{Left: n3, Right: n4, Val: 1}
	n0 := &TreeNode{Left: n1, Right: n2, Val: 0}

	// s2 := &TreeNode{Left: nil, Right: nil, Val: 2}
	// s3 := &TreeNode{Left: nil, Right: nil, Val: 3}
	// s4 := &TreeNode{Left: s2, Right: nil, Val: 4}
	// s1 := &TreeNode{Left: s3, Right: s4, Val: 1}
	// s0 := &TreeNode{Left: nil, Right: s1, Val: 0}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"测试1", args{n0}, 3},
		//{"测试2", args{s0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
