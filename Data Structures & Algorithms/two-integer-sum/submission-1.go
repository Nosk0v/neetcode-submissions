func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i := range nums {
		digit := target - nums[i]
		if v, ok := hmap[digit]; ok {
			return []int{v,i}
		}
		hmap[nums[i]] = i	
	}
	return nil
}