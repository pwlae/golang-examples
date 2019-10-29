package main

func mySqrt(x int) int {
    // right is a max sqrt of int type
    left, right := 0, 65535
    
    // mid of max int sqrt
    mid := 32768

    for right-left != 1 {
        if mid*mid == x {
            return mid
        }
        if mid*mid > x {
            right = mid
        } else {
            left = mid
        }

        mid = (left + right) / 2
    }
    return mid
}

func main() {
    print("result: ", mySqrt(49), "\n")
    print("result: ", mySqrt(16), "\n")
    print("result: ", mySqrt(24), "\n")
    print("result: ", mySqrt(25), "\n")
    print("result: ", mySqrt(112), "\n")
}
