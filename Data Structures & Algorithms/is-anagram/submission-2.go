func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var seen [26]int
	for i := range s {
		seen[s[i]-'a']++
	}

	for i := range t {
		seen[t[i]-'a']--
	}

	for i := range seen {
		if seen[i] != 0 {
			return false
		}
	}

	return true

}
