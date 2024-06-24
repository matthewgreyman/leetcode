// Runtime: 3 ms | Beats: 78.35% of golang submissions.
// Memory Usage: 3.9 MB | Beats: 11.47% of golang submissions.
func isAnagram(s string, t string) bool {
	m := make(map[string]int)
	m2 := make(map[string]int)
	for i := range s {
		m[string(s[i])] += 1
	}
	for j := range t {
		m2[string(t[j])] += 1
	}
	if reflect.DeepEqual(m, m2) {
		return true
	} else {
		return false
	}
}