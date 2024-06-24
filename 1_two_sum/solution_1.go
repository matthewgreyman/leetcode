// Runtime: 0ms | Beats: 100% of golang submissions.
// Memory Usage: 4.2 MB | Beats: 57.69% of golang submissions.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		needed := target - num

		if index, ok := m[needed]; ok {
			return []int{index, i}
		}
		m[num] = i
	}
	return []int{}
}