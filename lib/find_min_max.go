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
