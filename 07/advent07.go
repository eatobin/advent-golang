// Does this version correct:
// https://en.wikipedia.org/wiki/Heap%27s_algorithm#Frequent_mis-implementations

package main

import "fmt"

func permutations(k int, A []int) {
	if k == 1 {
		fmt.Println(A)
	} else {
		permutations(k-1, A)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				A[i], A[k-1] = A[k-1], A[i]
			} else {
				A[0], A[k-1] = A[k-1], A[0]
			}
			permutations(k-1, A)
		}
	}
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	permutations(len(A), A)
}
