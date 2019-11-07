// Good article about reader https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

/*
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
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
