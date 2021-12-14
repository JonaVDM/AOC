package lib

// copied from https://learningprogramming.net/golang/golang-golang/find-max-and-min-of-array-in-golang/
func FindMinMax(a [2]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func MinMaxList(nums []int) (int, int) {
	smallest := nums[0]
	biggest := nums[0]
	for _, n := range nums {
		if smallest > n {
			smallest = n
		}

		if biggest < n {
			biggest = n
		}
	}

	return smallest, biggest
}
