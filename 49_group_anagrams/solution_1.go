// Runtime: 8 ms | Beats: 99.71% of golang submissions.
// Memory Usage: 7.4 MB | Beats: 83.77% of golang submissions.
func groupAnagrams(strs []string) [][]string {
	var f [][]string
	m2 := make(map[string][]string)

	for i := range strs {
		m := hash(strs[i])
		m2[m] = append(m2[m], strs[i])
	}

	for _, value := range m2 {
		fmt.Println(value)
		f = append(f, value)
	}
	return f
}

func hash(s string) string {
	res := make([]byte, 26)
	for _, c := range s {
		res[c-'a'] += 1
	}
	return string(res)
}