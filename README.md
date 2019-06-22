# log

## benchstat

system: windows 7 x64 sp1 

cpu(s): 8 

model name: Intel(R) Core(TM) i7-3610QM CPU @ 2.30GHz 

memery: 8G 


### TextPositiveWithConsole

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkLogTextPositive-8       | 20000    | 76654 ns/op   | 4180 B/op   | 77 allocs/op    |

### TextPositiveWithoutConsole

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkLogTextPositive-8       | 200000   | 8005 ns/op    | 2482 B/op   | 48 allocs/op    |
