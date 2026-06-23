func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	var pairs []pair
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Freq > pairs[j].Freq
	})
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].Value)
	}
	return result
}

type pair struct {
	Value int
	Freq  int
}
