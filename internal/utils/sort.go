package utils

import "sort"

type Identifiable interface {
	Id() uint64
}

func SortById[T Identifiable](slice []T) []T {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Id() < slice[j].Id()
	})

	return slice
}
