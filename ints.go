package ints

import "math/rand"

// Max returns the maximum of a list of integers.
func Max(i ...int) int {
	switch len(i) {
	case 0:
		return 0
	case 1:
		return i[0]
	default:
		rest := Max(i[1:]...)
		if i[0] > rest {
			return i[0]
		}
		return rest
	}
}

// Min returns the minimum of a list of integers.
func Min(i ...int) int {
	switch len(i) {
	case 0:
		return 0
	case 1:
		return i[0]
	default:
		rest := Min(i[1:]...)
		if i[0] < rest {
			return i[0]
		}
		return rest
	}
}

// Shuffle shuffles the provided slice.
func Shuffle(i []int, r *rand.Rand) []int {
	l := len(i)
	if l <= 1 {
		return i
	}
	shuffled := make([]int, l)
	for k, v := range r.Perm(l) {
		shuffled[k] = i[v]
	}
	return shuffled
}

// Sum calculates the sum of values in a slice.
func Sum(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

// Find finds the first instance of a value in a slice.  Returns -1 when not
// found.
func Find(needle int, haystack []int) int {
	for k, v := range haystack {
		if v == needle {
			return k
		}
	}
	return -1
}

// Count counts the number of instances of a value in a slice.
func Count(needle int, haystack []int) int {
	count := 0
	for _, v := range haystack {
		if v == needle {
			count++
		}
	}
	return count
}
