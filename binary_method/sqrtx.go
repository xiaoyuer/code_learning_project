func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) >> 1
		if mid <= x/mid {
			l = mid
		} else {
			r = mid - 1
		}

	}
	return l
}

// func mySqrt(x int) int {
// 	res := x
// 	for res*res > x {
// 		res = (res + x/res) / 2
// 	}
// 	return res
// }