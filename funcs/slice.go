package funcs

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
