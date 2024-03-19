package funcs

import "reflect"

type SliceType interface{ comparable }

func SliceFilter[T SliceType](sliceToFilter []T, filterFunc func(item T) bool) []T {
	sliceAfterFilter := []T{}
	for _, eachItem := range sliceToFilter {
		if filterFunc(eachItem) {
			sliceAfterFilter = append(sliceAfterFilter, eachItem)
		}
	}

	return sliceAfterFilter
}

func SliceGetEnd[T any](list []T) T {
	return list[len(list)-1]
}

func Contains[E comparable](s []E, e E) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}
