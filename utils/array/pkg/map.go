package pkg

func Map[T any, R any](array []T, fn func(index int, item T) R) []R {
	var temp []R
	for index, item := range array {
		temp = append(temp, fn(index, item))
	}
	return temp
}
