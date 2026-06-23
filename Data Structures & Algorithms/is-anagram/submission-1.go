func isAnagram(s string, t string) bool {
	var chars [26]int
	for i := range s {
		chars[s[i]-'a']++
	}
	for i := range t {
		chars[t[i]-'a']--
	}

	for i := range chars {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}

