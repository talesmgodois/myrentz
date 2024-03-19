package lib

import "math/rand"

func Shuffle[T any](slice []T) []T {
	for i := range slice {
		rnd := rand.Intn(i + 1)
		slice[i], slice[rnd] = slice[rnd], slice[i]
	}
	return slice
}
