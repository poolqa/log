# log

## benchstat

system: windows 7 x64 sp1 

cpu(s): 8 

model name: Intel(R) Core(TM) i7-3610QM CPU @ 2.30GHz 

memery: 8G 


### TextPositiveWithConsole

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkLogTextPositive         | 30000    | 46702 ns/op   | 4016 B/op   | 78 allocs/op    |
| BenchmarkLogTextPositive-2       | 50000    | 24721 ns/op   | 4016 B/op   | 78 allocs/op    |
| BenchmarkLogTextPositive-4       | 100000   | 15950 ns/op   | 4017 B/op   | 78 allocs/op    |
| BenchmarkLogTextPositive-8       | 200000   | 11295 ns/op   | 4019 B/op   | 78 allocs/op    |

### TextPositiveWithoutConsole

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkLogTextPositive         | 100000   | 18831 ns/op   | 2480 B/op   | 48 allocs/op    |
| BenchmarkLogTextPositive-2       | 200000   | 10010 ns/op   | 2480 B/op   | 48 allocs/op    |
| BenchmarkLogTextPositive-4       | 200000   |  6840 ns/op   | 2481 B/op   | 48 allocs/op    |
| BenchmarkLogTextPositive-8       | 200000   |  7905 ns/op   | 2482 B/op   | 48 allocs/op    |
