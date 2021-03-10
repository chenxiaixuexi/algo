package ics4

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	hashmap := map[rune]int{}
	result := 0
	start := 0
	for i, j := range s {
		if _, ok := hashmap[j]; ok {
			//abba不能回到a每次都要取max
			start = Max(start, hashmap[j]+1)
		}
		result = Max(result, i+1-start)
		hashmap[j] = i
	}
	return result
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
