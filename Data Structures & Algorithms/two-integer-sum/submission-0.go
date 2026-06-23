func twoSum(nums []int, target int) []int {
	if len(nums) < 2 || len(nums) > 1000 {
		return nil
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}