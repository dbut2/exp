package math

import (
	"sort"
)

func Max[N number](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func Min[N number](a, b N) N {
	if a < b {
		return a
	}
	return b
}

func Abs[N number](a N) N {
	if a < 0 {
		return -a
	}
	return a
}

func Sign[N number](a N) N {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func Order[T number](s []T, desc bool) []T {
	t := s
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	if desc {
		t = Reverse(t)
	}
	return t
}

func OrderMap[T any, N number](s []T, f func(T) N, desc bool) []T {
	t := s
	sort.Slice(s, func(i, j int) bool {
		return f(t[i]) < f(t[j])
	})
	if desc {
		t = Reverse(t)
	}
	return t
}

func LargestN[N number](s []N, n int) []N {
	if n <= 0 {
		return []N{}
	}

	if n >= len(s) {
		return Order(s, true)
	}

	l := s[:n]

	mini := 0
	for i := 1; i < len(l); i++ {
		if l[i] < l[mini] {
			mini = i
		}
	}

	for i := n; i < len(s); i++ {
		if s[i] <= l[mini] {
			continue
		}

		s[mini] = s[i]
		for j := range l {
			if l[j] < l[mini] {
				mini = j
			}
		}
	}

	return Order(l, true)
}

func Largest[N number](s []N) N {
	return LargestN(s, 1)[0]
}

func LargestNMap[T any, N number](s []T, f func(T) N, n int) []T {
	return OrderMap(s, f, true)[:n]
}

func LargestMap[T any, N number](s []T, f func(T) N) T {
	return LargestNMap(s, f, 1)[0]
}

func SmallestN[N number](s []N, n int) []N {
	if n <= 0 {
		return []N{}
	}

	if n >= len(s) {
		return Order(s, false)
	}

	l := s[:n]

	maxi := 0
	for i := 1; i < len(l); i++ {
		if l[i] > l[maxi] {
			maxi = i
		}
	}

	for i := n; i < len(s); i++ {
		if s[i] >= l[maxi] {
			continue
		}

		s[maxi] = s[i]
		for j := range l {
			if l[j] > l[maxi] {
				maxi = j
			}
		}
	}

	return Order(l, false)
}

func Smallest[N number](s []N) N {
	return SmallestN(s, 1)[0]
}

func SmallestNMap[T any, N number](s []T, f func(T) N, n int) []T {
	return OrderMap(s, f, false)[:n]
}

func SmallestMap[T any, N number](s []T, f func(T) N) T {
	return SmallestNMap(s, f, 1)[0]
}

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T number](s []T) T {
	t := T(0)
	for _, i := range s {
		t += i
	}
	return t
}

func SumMap[T comparable, N number](s map[T]N) N {
	t := N(0)
	for _, i := range s {
		t += i
	}
	return t
}

func SumMapIf[T comparable, N number](s map[T]N, predicate func(T) bool) N {
	t := N(0)
	for k, v := range s {
		if predicate(k) {
			t += v
		}
	}
	return t
}

func Pow[N number](x, y N) N {
	val := N(1)
	for i := N(0); i < y; i++ {
		val *= x
	}
	return val
}
