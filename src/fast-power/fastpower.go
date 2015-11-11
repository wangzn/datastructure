package fastpower

import ()

//Calculate the an % b where a, b and n are all 32bit integers.
func FastPower(a, b, n int) int {
	if n == 1 {
		return a % b
	}
	if n == 0 {
		return 1 % b
	}
	if n < 0 {
		return -1
	}
	p := FastPower(a, b, n/2)
	p = (p * p) % b
	if n%2 == 1 {
		p = (p * a) % b
	}
	return p
}
