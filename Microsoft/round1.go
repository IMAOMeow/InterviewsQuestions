// https://leetcode.com/problems/jump-game-ii/
// written during interview, can't promise the correctness, just for draft reference

var minJump int = math.MaxInt

// Additionally: use map to store calcuated value -> o(n) time complexity
var indexToMinJump = make(map[int]int)

func jump(nums []int) int {
	// // special cases
	// // 1. len = 0  2. len>=1, nums[0] == 0
	// if len(nums) == 0 {
	// 	return 0
	// }
	search(nums, 0, 0)
	return minJump
}

func search(nums []int, index, count int) {
	// index valid
	if index < 0 || index >= len(nums) {
		return
	}

	// end condition
	if index == len(nums)-1 {
		if count < minJump {
			minJump = count
		}

		return
	}

	// jump
	times := nums[index]

	count++
	for i := times; i >= 1; i-- {
		search(nums, index+i, count)
	}
}
