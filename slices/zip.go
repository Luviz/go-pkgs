package slices

type Pair[T, U any] struct {
	First  T
	Second U
}

func Zip[T, U any](s1 []T, s2 []U) []Pair[T, U] {
	res := make([]Pair[T, U], 0)
	min_len := min(len(s1), len(s2))
	for i := 0; i < min_len; i++ {
		res = append(res, Pair[T, U]{s1[i], s2[i]})
	}
	return res
}
