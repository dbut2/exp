package slices

func Fold[T, U any](s []T, initial U, f func(U, T) U) U {
	for _, v := range s {
		initial = f(initial, v)
	}
	return initial
}
