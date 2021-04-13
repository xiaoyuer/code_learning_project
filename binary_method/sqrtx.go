package main

import "fmt"

// func mySqrt(x int) int {
// 	l, r := 0, x
// 	for l < r {
// 		mid := (l + r + 1) >> 1
// 		if mid <= x/mid {
// 			l = mid
// 		} else {
// 			r = mid - 1
// 		}

// 	}
// 	return l
// }
func main() {
	mySqrt(8)
}

func mySqrt(x float64) float64 {
	res := x
	for res*res > x {
		fmt.Println(res)
		res = (res + x/res) / 2
	}
	return res
}

// func main() {
//     fmt.Println(sqrt(3))

// }

// func sqrt(x float64)float64{
//      z := x
//     for i := 0; i < 10 ; i++  {
//         z = z - (z*z -x)/(2*z)
//     }
//     return z
// }
