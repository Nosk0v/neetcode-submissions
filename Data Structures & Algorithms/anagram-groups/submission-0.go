func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var chars [26]int
		for _, ch := range str {
			chars[ch-'a']++
		}
		groups[chars] = append(groups[chars],str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
