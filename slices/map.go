package slices

func Map[T, U any](s []T, cb func(item T, index int, slice []T) U) []U {
	res := make([]U, 0)

	for ix, v := range s {
		res = append(res, cb(v, ix, s))
	}

	return res
}
