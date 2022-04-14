package main

import (
	"AlgorithmsFourthEdition/Chapter02Sorting"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		intn := rand.Intn(100)
		a = append(a, intn)
	}
	fmt.Println(a)
	Chapter02Sorting.MergeSortBottom2Top(a)
	fmt.Println(a)
}
