package bsearch

import ()

func Search(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == t {
			return 0
		} else {
			return -1
		}
	}
	st := 0
	ed := len(a)
	for st < ed {
		mid := (st + ed) / 2
		if a[mid] == t {
			return mid
		} else if t < a[mid] {
			ed = mid - 1
		} else {
			st = mid + 1
		}
	}
	return -1
}
