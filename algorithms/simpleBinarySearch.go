package main

func search(nums []int, target int) int {
    if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
	}

	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{1,2,3,4,5,6,7,8,9,99,101}
	print("position: ", search(list,2), "\n")
	print("position: ", search(list,3), "\n")
	print("position: ", search(list,4), "\n")
	print("position: ", search(list,5), "\n")
	print("position: ", search(list,6), "\n")
	print("position: ", search(list,7), "\n")
	print("position: ", search(list,72348723894), "\n")
	print("position: ", search(list,101), "\n")
}
