package fyshuffle

import (
	"math/rand"
	"time"
)

func FYShuffle(a []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}
