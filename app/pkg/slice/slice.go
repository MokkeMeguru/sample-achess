package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
	"golang.org/x/xerrors"
)

func Map[T, V any](s []T, fn func(T) V) []V {
	rs := make([]V, len(s))
	for idx, e := range s {
		rs[idx] = fn(e)
	}
	return rs
}

func MapWithError[T, V any](s []T, fn func(T) (V, error)) ([]V, error) {
	rs := make([]V, len(s))
	for idx, element := range s {
		r, err := fn(element)
		if err != nil {
			return nil, err
		}
		rs[idx] = r
	}
	return rs, nil
}

func Copy[T any](s []T) []T {
	rs := make([]T, len(s))
	copy(s, rs)
	return rs
}

func Max[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, xerrors.New("unexpected array: len(s) == 0")
	}
	m := s[0]
	for _, element := range s {
		if m < element {
			m = element
		}
	}
	return m, nil
}

func Min[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, xerrors.New("unexpected array: len(s) == 0")
	}
	m := s[0]
	for _, element := range s {
		if m < element {
			m = element
		}
	}
	return m, nil
}

func ArrayToMap[T any, K constraints.Ordered](s []T, keyFn func(T) K) map[K]T {
	rm := make(map[K]T, len(s))
	for _, e := range s {
		rm[keyFn(e)] = e
	}
	return rm
}

func ArrayToMapWithError[T any, K constraints.Ordered](s []T, keyFn func(T) (K, error)) (map[K]T, error) {
	rm := make(map[K]T, len(s))
	for _, e := range s {
		k, err := keyFn(e)
		if err != nil {
			return nil, err
		}
		rm[k] = e
	}
	return rm, nil
}

func Count[T any](s []T, compareFn func(T) bool) int {
	r := 0
	for _, e := range s {
		if compareFn(e) {
			r++
		}
	}
	return r
}

func CopySort[T any](s []T, less func(a, b T) bool) []T {
	rs := Copy(s)
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return rs
}

func CopySortStable[T any](s []T, less func(a, b T) bool) []T {
	rs := Copy(s)
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return rs
}

func Contains[T comparable](s []T, target T) bool {
	for _, e := range s {
		if e == target {
			return true
		}
	}
	return false
}

func Has[T any](s []T, hasFn func(s T) bool) bool {
	for _, e := range s {
		if hasFn(e) {
			return true
		}
	}
	return false
}
