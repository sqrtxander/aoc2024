package utils

import "slices"

func Filter[T any](slice []T, test func(T) bool) (result []T) {
	for _, s := range slice {
		if test(s) {
			result = append(result, s)
		}
	}
	return
}

func RemoveAll[T comparable](slice []T, object T) (result []T) {
	result = Filter(slice, func(o T) bool {
		return object != o
	})
	return
}

func AreSetEqual[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	if !All(slice1, func(o T) bool {
		return slices.Contains(slice2, o)
	}) {
		return false
	}
	return All(slice2, func(o T) bool {
		return slices.Contains(slice1, o)
	})
}

func RemoveFirst[T comparable](slice []T, object T) (result []T) {
	found := false
	for _, s := range slice {
		if found || s != object {
			result = append(result, s)
		} else if s == object {
			found = true
		}
	}
	return
}

func Map[T any, S any](slice []T, change func(T) S) (result []S) {
	for _, s := range slice {
		result = append(result, change(s))
	}
	return
}

func Any[T any](slice []T, test func(T) bool) bool {
	for _, s := range slice {
		if test(s) {
			return true
		}
	}
	return false
}

func All[T any](slice []T, test func(T) bool) bool {
	for _, s := range slice {
		if !test(s) {
			return false
		}
	}
	return true
}
