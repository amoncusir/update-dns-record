package arrays

func OnMap[T, V any](arr []T, f func(T) V) []V {
	acc := make([]V, len(arr))

	for i := range arr {
		acc[i] = f(arr[i])
	}

	return acc
}
