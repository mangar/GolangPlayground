# <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" width="70px"> Golang Playground



## Test 1
Testing and __Benchmark__ about verify if a key exists in a map or just add w/out checking.
Method `RepeatCount` is faster than `RepeatCountIfExist`

__Result__

```go test -bench=.```


    cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
    BenchmarkRepeatCount-8                  1000000000               0.0000056 ns/op               0 B/op          0 allocs/op
    BenchmarkRepeatCountIFExist-8           1000000000               0.0000082 ns/op               0 B/op          0 allocs/op


