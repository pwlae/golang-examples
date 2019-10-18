package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	mindex, nindex := 0, 0
	result := []int{}

	for mindex < m && nindex < n {
		if nums1[mindex] > nums2[nindex] {
			result = append(result, nums2[nindex])
			nindex += 1
			continue
		}
		result = append(result, nums1[mindex])
		mindex += 1
	}

	for mindex < m {
		result = append(result, nums1[mindex])
		mindex += 1
	}

	for nindex < n {
		result = append(result, nums2[nindex])
		nindex += 1
	}


	for index,value := range result {
		nums1[index] = value
		//print(nums1[index])
	}
}

func main() {
	array1 := []int{1, 2, 3, 0, 0, 0}
	array2 := []int{2, 5, 6}

	merge(array1, 3, array2, 3)
}
