package slices

func Filter[T comparable](slice []T, cb func(item T, ix int, slice []T) bool) []T {
	res := make([]T, 0)
	for ix, v := range slice {
		if cb(v, ix, slice) {
			res = append(res, v)
		}
	}
	return res
}
