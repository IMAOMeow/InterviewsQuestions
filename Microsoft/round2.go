// https://leetcode.com/problems/sort-colors
// written during interview, can't promise the correctness, just for draft reference

// sort
// [2,0,2,1,1,0] -> [0,0,1,1,2,2]
// Method 1: classify through counting for each number, then appending

func sortColors (nums []int){
	numToCount := make([]int, 3)

	for _, num := range nums {
		numToCount[num]++
	}

	index := 0
	for i, v := range numToCount {
		for j:=0; j<v; j++ {
			nums[index] = i
			index++
		}
	}
}

// sort
// [2,0,2,1,1,0] -> [0,0,1,1,2,2]
// Method 2: two pointers, simulating quick sort, each time group one color, need to go through the process two times

func sortColors (nums []int){
	// input validation ?
	if len(nums) < 2 {
		return
	}

	// sort
	left, right := 0, 1
	pivot := nums[0]
	
	//end condition: right - > end 
	for right < len(nums) {
		for nums[left]==pivot {
			left++
		}
	
		for nums[right] != pivot {
			right++
		}
	
		// switch
		nums[left], nums[right] = nums[right], nums[left]
	}

	// iteration times control? 3->2
	sortColors(nums[left:])
}

// Method 3: iterate with two pointers, left:=0, right := len(nums), then iterate, if nums[i] == 0, switch with left, if nunms[i] == 2, switch with right
