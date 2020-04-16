## RangeCount
Consider the problem of counting occurrences of a given value x within a sorted array a.
For example, if x=5 and a=[1,1,2,4,5,5,7,9], then the count is 2.

I. Write a function that solves this problem by performing a linear scan.
`func linearScan(...) ...`

Ii. Next, write a function that solves this problem by performing two binary searches.
`func twoBinarySearchScan(...)...`

### Benchmark
Benchmark both the functions for random sorted arrays of size 10, 100, ..., up to 10,000,000.
```
$ go test -bench=.
goos: darwin
goarch: amd64

BenchmarkLinearScan10-12                    	1000000000	         0.000029 ns/op
BenchmarkLinearScan100-12                   	1000000000	         0.000034 ns/op
BenchmarkLinearScan1000-12                  	1000000000	         0.000159 ns/op
BenchmarkLinearScan10000-12                 	1000000000	         0.00136 ns/op
BenchmarkLinearScan100000-12                	1000000000	         0.0168 ns/op
BenchmarkLinearScan1000000-12               	1000000000	         0.192 ns/op
BenchmarkLinearScan10000000-12              	       1	           2274060143 ns/op
BenchmarkLinearScan100000000-12             	       1	           25477415442 ns/op

BenchmarkTwoBinarySearchScan10-12           	1000000000	         0.000010 ns/op
BenchmarkTwoBinarySearchScan100-12          	1000000000	         0.000079 ns/op
BenchmarkTwoBinarySearchScan1000-12         	1000000000	         0.000115 ns/op
BenchmarkTwoBinarySearchScan10000-12        	1000000000	         0.00133 ns/op
BenchmarkTwoBinarySearchScan100000-12       	1000000000	         0.0195 ns/op
BenchmarkTwoBinarySearchScan1000000-12      	1000000000	         0.192 ns/op
BenchmarkTwoBinarySearchScan10000000-12     	       1	           2246179923 ns/op
BenchmarkTwoBinarySearchScan100000000-12    	       1	           25030702403 ns/op

PASS
ok  	_/Users/random/serene-garden-57635/range_count	61.067s
```

How does performance compare between the two functions?
By looking at the Benchmark data, using twoBinarySearchScan seems to be quicker solution than linearScan.
