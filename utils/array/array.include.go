package array

func Include[T string | int](array []T, val T) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}
