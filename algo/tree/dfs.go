package ics4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	//res := []*TreeNode{}
	num := 0
	counter := 0
	//left := 0
	//right := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		//res = append(res, node)
		queue = queue[:len(queue)-1]
		if node.Right != nil || node.Left != nil {
			counter++
		} else {
			num = Max(num, counter+1)
			//判断是否回到了上一层
			//counter--
			//去了左边节点或者右边节点就不能--
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			//right++
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			//left++
		}
	}

	return num
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
