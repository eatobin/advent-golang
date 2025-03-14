package main

import "fmt"

func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	} else {
		for i := 0; i < size-1; i++ {
			HeapPermutation(a, size-1)
			if size%2 == 0 {
				a[i], a[size-1] = a[size-1], a[i]
			} else {
				a[0], a[size-1] = a[size-1], a[0]
			}
		}
		HeapPermutation(a, size-1)
	}
}

func main() {
	a := []int{1, 2, 3}
	HeapPermutation(a, len(a))
}
