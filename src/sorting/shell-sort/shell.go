package shell

import (
	"fmt"
)

func ShellSort(a []int) []int {
	var increment int
	increment = len(a) / 2
	for increment > 0 {
		for i := 1; i < len(a); i++ {
			j := i - increment
			tmp := a[i]
			for j >= 0 && tmp < a[j] {
				a[j+increment] = a[j]
				j = j - increment
			}
			a[j+increment] = tmp
		}
		fmt.Println(a)
		if increment == 2 {
			increment = 1
		} else {
			increment = increment / 2
		}
	}
	return a
}
