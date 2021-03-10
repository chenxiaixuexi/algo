package ics4

func Sqrt(x int) int {
	left := 1
	right := x / 2
	mid := (left + right) / 2
	for left < right {
		//mid = (left + right) / 2
		//???? 增加搜索范围
		mid = (right+left)/2 + 1
		if mid*mid > x {
			//只有大于才减
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
