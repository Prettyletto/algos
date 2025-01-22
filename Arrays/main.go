package main

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

func LinearSearch[T comparable](haystack []T, needle T) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func BinarySearch[T Ordered](haystack []T, low, high int, needle T) bool {

	if low > high {
		return false
	}
	middle := low + (high-low)/2

	if haystack[middle] == needle {
		return true
	}

	if haystack[middle] < needle {
		return BinarySearch(haystack, middle+1, high, needle)

	}
	return BinarySearch(haystack, low, middle, needle)

}

func BubbleSort[T Ordered](haystack []T) {
	for i := range haystack {
		for k := 0; k < len(haystack)-1-i; k++ {
			if haystack[k] > haystack[k+1] {
				tmp := haystack[k]
				haystack[k] = haystack[k+1]
				haystack[k+1] = tmp
			}
		}
	}
}

func main() {

	x := []int{1, 2, 32, 53, 4, 5, 29, 18, 23, 1212, 222, 1010}
	fmt.Println(x)
	BubbleSort(x)
	fmt.Println(x)

}
