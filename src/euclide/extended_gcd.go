package euclide

import ()

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	} else if a < b {
		return Gcd(b-a, a)
	} else {
		return Gcd(a-b, b)
	}
}

func ExtGcd(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	tx, ty := ExtGcd(b, a%b)
	return ty, tx - a/b*ty
}
