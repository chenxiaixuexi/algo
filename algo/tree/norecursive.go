package ics4

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth_new(node *TreeNode) int {
	result := 0
	// var stack []*TreeNode
	// lastVisit := node
	// for {
	// 	for node != nil {
	// 		stack = append(stack, node)
	// 		node = node.Left
	// 	}
	// 	if len(stack) == 0 {
	// 		break
	// 	}
	// 	node = stack[len(stack)-1]
	// 	//保证不会访问第二次陷入循环
	// 	if node.Right != nil && node.Right != lastVisit {
	// 		//首次访问
	// 		node = node.Right
	// 		continue
	// 	}
	// 	//后序遍历来统计层级
	// 	if result < len(stack) {
	// 		result = len(stack)
	// 	}
	// 	stack = stack[:len(stack)-1]
	// 	lastVisit = node
	// 	node = nil
	// }

	//lll
	stack := []*TreeNode{}
	depth := []int{1}

	if node == nil {
		return 0
	}
	stack = append(stack, node)
	for len(stack) != 0 {
		thisDepth := depth[len(depth)-1]
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		depth = depth[:len(depth)-1]
		if node == nil {
			continue
		}
		if result < thisDepth {
			result = thisDepth
		}
		stack = append(stack, node.Left, node.Right)
		depth = append(depth, thisDepth+1, thisDepth+1)
	}

	return result

}
func maxDepthBfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	stack_two := []*TreeNode{}
	res := [][]int{}
	level := []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		for len(stack) != 0 {
			node := stack[len(stack)-1]
			stack_two = append(stack_two, node)
			stack = stack[:len(stack)-1]
			if node == nil {
				continue
			}
			level = append(level, node.Val)
		}
		for len(stack_two) != 0 {
			node := stack_two[len(stack_two)-1]
			stack_two = stack_two[:len(stack_two)-1]
			if node == nil {
				continue
			}
			stack = append(stack, node.Right, node.Left)
		}
		if len(level) != 0 {
			res = append(res, level)
			level = []int{}
		}
	}
	return res
}
