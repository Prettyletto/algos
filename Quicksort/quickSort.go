package main

import "fmt"

type Ordered interface {
	~int | string | ~float64
}

func partition[T Ordered](arr []T, low, high int) int {
	pivot := arr[high]

	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}
	idx++
	arr[high] = arr[idx]
	arr[idx] = pivot
	return idx
}

func qs[T Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}
	pivotIdx := partition(arr, low, high)

	qs(arr, low, pivotIdx-1)
	qs(arr, pivotIdx+1, high)

}

func QuickSort[T Ordered](arr []T) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	x := []int{0, 22, 33, 1, 2, 8, 98, 22, 3, 7, 198, 9, 8, 11}
	QuickSort(x)
	fmt.Println(x)

}
