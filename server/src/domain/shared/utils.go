package shared

import "slices"

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i, v := range ts {
		us[i] = f(v)
	}
	return us
}

func Filter[T any](s []T, f func(el T) bool) []T {
	return slices.DeleteFunc(slices.Clone(s), func(el T) bool { return !f(el) })
}
