package lib

func Max(arr []int) (index int, value int) {
	var max int
	for idx, val := range arr {
		if val > max {
			max = val
			index = idx
		}
	}
	return index, max
}
