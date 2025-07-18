package main

import (
	"fmt"
	"math/rand"
	TDAPQ "tdas/cola_prioridad"
)

func cmp(a, b int) int {
	return a - b
}

func TopK(arr []int, k int) []int {
	cp := TDAPQ.CrearHeapArr(arr, cmp)
	top := make([]int, k)

	for i := 0; i < k; i++ {
		top[i] = cp.Desencolar()
	}

	return top
}

func TopKStream(arr []int, k int) []int {
	cp := TDAPQ.CrearHeapArr(arr[:k], func(a, b int) int { return cmp(b, a) })
	for _, elem := range arr[k:] {
		if elem > cp.VerMax() {
			cp.Desencolar()
			cp.Encolar(elem)
		}
	}

	top := make([]int, k)
	for i := 0; !cp.EstaVacia(); i++ {
		top[k-i-1] = cp.Desencolar()
	}
	return top
}

func main() {
	miarr := make([]int, 10000)
	for i := range miarr {
		miarr[i] = rand.Intn(50000)
	}
	fmt.Println(TopK(miarr, 10))
	fmt.Println(TopKStream(miarr, 10))

}
