// Runtime: 4 ms | Beats: 97.68% of golang submissions.
// Memory Usage: 6.32 MB | Beats: 34.15% of golang submissions.
func topKFrequent(nums []int, k int) []int {
	var res []int
	m := make(map[int]int, len(nums))
	m2 := make(map[int][]int, len(nums))
	for i := range nums {
		m[nums[i]]++
	}
	for number, times := range m {
		m2[times] = append(m2[times], number)
	}
	for j := len(nums) + 1; j > 0; j-- {
		for l := range m2[j] {
			res = append(res, m2[j][l])
			if len(res) == k {
				return res
			}
		}
	}
	return res
}