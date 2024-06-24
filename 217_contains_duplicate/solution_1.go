// Runtime: 69 ms | Beats: 91.09% of golang submissions.
// Memory Usage: 8.9 MB | Beats: 69.86% of golang submissions.
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] == true {
			return true
		}
		m[nums[i]] = true
	}
	return false
}