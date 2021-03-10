package ics4

func FindMin(nums []int) int {
	//consider [2,1]
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	// 	if nums[left] > nums[right] || len(nums) == 1 {
	// 		for left < right {
	// 			mid = (left + right) / 2
	// 			if nums[left] < nums[mid] {
	// 				left = mid - 1
	// 			} else if nums[left] > nums[mid] {
	// 				right = mid + 1
	// 			}
	// 		}
	// 	} else {
	// 		return nums[left]
	// 	}
	// 	return nums[mid]
	// }

	if nums[left] > nums[right] || len(nums) != 1 {
		for left < right {
			mid = (left + right) / 2
			if mid > 0 && nums[mid] < nums[mid-1] {
				return nums[mid]
			}
			if nums[left] <= nums[mid] && nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	} else {
		return nums[left]
	}
	return nums[left]
}
