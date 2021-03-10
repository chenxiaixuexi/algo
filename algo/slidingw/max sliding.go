package ics4

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := []int{}
	result := []int{}
	for i := range nums {
		//&& 先要判断queue是否为空
		for i > 0 && len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			//去尾
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		//亮点，查看k个之前有没有被删掉，无论是相等还是单调递减都可以删掉，如果递增已经被删掉
		if i >= k && nums[i-k] == queue[0] {
			//tou
			queue = queue[1:]
		}
		// if len(queue) > k {
		// 	queue = queue[1:]
		// }
		// 会发生一些元素超过窗口生命周期还是存在
		if i >= k-1 {
			result = append(result, queue[0])
		}

	}
	return result
}
