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

func SliceContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}

func SliceRemoveDuplicateElement[T SliceType](sliceToRemoveDuplicate []T) []T {
	result := make([]T, 0, len(sliceToRemoveDuplicate))
	set := map[T]struct{}{}
	for _, item := range sliceToRemoveDuplicate {
		if _, ok := set[item]; !ok {
			set[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
