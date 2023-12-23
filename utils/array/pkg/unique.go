package pkg

func Unique[T any](array []T, fn func(index int, item T) string) []T {
	set := make(map[string]T, len(array))
	flag := 0
	var key string

	for index, value := range array {
		key = fn(index, value)
		if _, ok := set[key]; !ok {
			set[key] = value
			array[flag] = value
			flag++
		}
	}

	return array
}
