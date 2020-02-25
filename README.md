# expensive-api

expensive-api is a dummy http API that uses CPU resources on purpose.
It is used to do CPU benchmarks (cfs quotas, cpuset, etc.)

## Commands

run it localy:
`AUTOMAXPROC=true go run main.go`

run it in docker with a CPU quota and go.uber.org/automaxprocs enabled
`docker run --cpu-quota=200000 -e AUTOMAXPROC=true  -p 8080:8080 7felf/expensive-api`

run a benchmark in docker
`docker run --net=host -v $PWD/benchmark.yml:/benchmark.yml xridge/drill --benchmark /benchmark.yml -s`

## Example testing the difference with and without go.uber.org/automaxprocs enabled

 *cpu-quota=200000 AUTOMAXPROC=false:*
 ```
Concurrency Level         8
Time taken for tests      35.1 seconds
Total requests            400
Successful requests       400
Failed requests           0
Requests per second       11.40 [#/sec]
Median time per request   659ms
Average time per request  668ms
Sample standard deviation 158ms
 ```
 *cpu-quota=200000 AUTOMAXPROC=true:*
 ```
 Concurrency Level         8
Time taken for tests      30.6 seconds
Total requests            400
Successful requests       400
Failed requests           0
Requests per second       13.08 [#/sec]
Median time per request   573ms
Average time per request  581ms
Sample standard deviation 145ms
 ```
