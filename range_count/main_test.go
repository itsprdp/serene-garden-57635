package main

import (
	"testing"
)

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
