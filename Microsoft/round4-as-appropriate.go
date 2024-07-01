// sorted []int, target int
// return index/-1
func binarySearch(arr []int, target, i, j int) int {
	// input validation
	if i > j || i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		return -1
	}

	mid := (i + j) / 2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return binarySearch(arr, i, mid-1)
	}

	return binarySearch(arr, mid+1, j)
}

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return
	}

	i, j := 0, len(arr)-1

	for i < j {
		mid := (i + j) / 2

		if arr[mid] == target {
			break
		} else if arr[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	if arr[i] == target {
		return i
	}

	return -1
}

// 1. input arr = []int{}, target = any -> -1
// 2. arr := {1, 2, 3, 4 ,5 } target =4 -> 3
// 3. arr := {1, 2, 3, 4 ,5 } target is not existed  -1, -> -1
// 4. arr := [ 1 , 1 , 1 , 1, 1] 
