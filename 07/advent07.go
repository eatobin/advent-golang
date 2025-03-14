package main

import "fmt"

func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	} else {
		HeapPermutation(a, size-1)
		for i := 0; i < size-1; i++ {
			if size%2 == 0 {
				a[i], a[size-1] = a[size-1], a[i]
			} else {
				a[0], a[size-1] = a[size-1], a[0]
			}
			HeapPermutation(a, size-1)
		}
	}
}

func main() {
	a := []int{1, 2, 3}
	HeapPermutation(a, len(a))
}
