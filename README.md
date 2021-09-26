![Build Status](https://github.com/StuartsHome/carPedia/actions/workflows/go.yml/badge.svg)  
# wikipedia for cars

## How to use
### How to run
#### Local
- `go run main.go` for production
- `go run main.go -development=true` for development "testing"

#### Container
- `make docker-start`


#### Logging
- logs are located in log/*

## Web Pages
- **/home**: web page that provides html form for saving car details into db
- **/allcars**: displays all the cars in the database in a html table
- **/car**: displays all the cars in the database in JSON format

## Features
- MySQL Database
- Makefile w/ make commands
- Containerised
- Static web page (HTML, JS)
- Shell scripts
- Logging
- Benchmarking and instrumentation via go-wrk and pprof


## Benchmarks
#### /allcars
```shell
==========================BENCHMARK==========================
URL:                            http://localhost:8100/allcars

Used Connections:               100
Used Threads:                   1
Total number of calls:          1000

===========================TIMINGS===========================
Total time passed:              2.78s
Avg time per request:           252.54ms
Requests per second:            359.69
Median time per request:        232.13ms
99th percentile time:           702.57ms
Slowest time for request:       736.00ms

=============================DATA=============================
Total response body sizes:              42114000
Avg response body per request:          42114.00 Byte
Transfer rate per second:               15148168.90 Byte/s (15.15 MByte/s)
==========================RESPONSES==========================
20X Responses:          1000    (100.00%)
30X Responses:          0       (0.00%)
40X Responses:          0       (0.00%)
50X Responses:          0       (0.00%)
Errors:                 0       (0.00%)
```


## Examples
Populate cache
```
curl -X POST -d '@exampleRequests/populateCache.json' http://localhost:8100/desc
```

Get cache descs
```
curl -X GET http://localhost:8100/desc
```