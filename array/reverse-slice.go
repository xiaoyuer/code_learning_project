package main

import "fmt"

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	Reverse(s)
	// sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

}

func Reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
