package main

import "fmt"

func main() {
	SliceToArray()
	Generic()
}

func SliceToArray() {
	// slice to array
	sl := []int{1, 2, 3, 4}
	ar1 := [2]int(sl)
	ar2 := [4]int(sl)
	fmt.Printf("slice(len=4) to array(len=2): %v\n", ar1)
	fmt.Printf("slice(len=4) to array(len=4): %v\n", ar2)
	ar2[1] = 3
	fmt.Printf("slice(len=4): %v\n", sl)
	fmt.Printf("array[1]=3: %v\n", ar1)
	fmt.Printf("array(no change): %v\n", ar2)
}

func Generic() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  1.23,
		"second": 4.56,
	}
	fmt.Printf("Generic sums: int=%v, float=%v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats),
	)
	fmt.Printf("Generic sums type inferred: int=%v, float=%v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
	fmt.Printf("Generic sums with interface constraint: int=%v, float=%v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 1.21
//func ClearMap() {
//	m := map[string]int{"first": 1, "second": 2}
//	p := m
//	clear(m)
//	fmt.Println(m)
//	fmt.Println(p)
//}
