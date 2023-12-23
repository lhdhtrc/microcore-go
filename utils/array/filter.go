package array

func Filter[T any](array []T, fn func(index int, item T) bool) []T {
	var temp []T
	for index, item := range array {
		if fn(index, item) {
			temp = append(temp, item)
		}
	}
	return temp
}
