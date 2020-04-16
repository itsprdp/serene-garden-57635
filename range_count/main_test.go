package main

import (
	"math/rand"
	"sort"
	"testing"
)

// custom []int32 to implement sort.Sort interface
type int32arr []int32

func (f int32arr) Len() int {
	return len(f)
}

func (f int32arr) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f int32arr) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// generate random sorted integers array
func randInt32Array(n int32) int32arr {
	elements := int32arr{}

	r := rand.New(rand.NewSource(99))

	for i := int32(0); i < n; i++ {
		// generate random index
		w := r.Int31n(n)
		elements = append(elements, w)
	}

	sort.Sort(elements)
	return elements
}

func BenchmarkLinearScan10(b *testing.B) {
	i := int32(10)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan100(b *testing.B) {
	i := int32(100)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan1000(b *testing.B) {
	i := int32(1000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan10000(b *testing.B) {
	i := int32(10000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan100000(b *testing.B) {
	i := int32(100000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan1000000(b *testing.B) {
	i := int32(1000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan10000000(b *testing.B) {
	i := int32(10000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkLinearScan100000000(b *testing.B) {
	i := int32(100000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	linearScan(arr, val)
}

func BenchmarkTwoBinarySearchScan10(b *testing.B) {
	i := int32(10)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan100(b *testing.B) {
	i := int32(100)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan1000(b *testing.B) {
	i := int32(1000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan10000(b *testing.B) {
	i := int32(10000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan100000(b *testing.B) {
	i := int32(100000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan1000000(b *testing.B) {
	i := int32(1000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan10000000(b *testing.B) {
	i := int32(10000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}

func BenchmarkTwoBinarySearchScan100000000(b *testing.B) {
	i := int32(100000000)
	arr := randInt32Array(i)
	val := arr[int32(i/2+1)]
	twoBinarySearchScan(arr, val)
}
