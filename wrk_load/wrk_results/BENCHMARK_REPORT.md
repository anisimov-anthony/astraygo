# WRK Benchmark Report - AstrayGo

**Generated:** 2026-01-29 00:27:17 MSK

## System Specifications

**CPU:**
Model name:                              Intel(R) Core(TM) i3-7020U CPU @ 2.30GHz
- Cores: 4
- Threads: 4

**Memory:**
Total: 7.2Gi, Available: 2.0Gi

**Operating System:**
Ubuntu 25.10

**Kernel:**
6.17.0-8-generic

**Architecture:**
x86_64

---

## 01_warmup_post


```
Running 10s test @ http://localhost:8080
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.96ms    1.05ms  24.06ms   90.87%
    Req/Sec   403.24     43.67   717.00     93.28%
  16145 requests in 10.10s, 3.53MB read
Requests/sec:   1599.00
Transfer/sec:    357.95KB
```

---

## 02_read_by_id


```
Running 10s test @ http://localhost:8080
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.10ms  744.95us  13.15ms   86.86%
    Req/Sec     1.87k   340.16     2.55k    66.25%
  74347 requests in 10.01s, 14.62MB read
Requests/sec:   7426.91
Transfer/sec:      1.46MB
```

---

## 03_read_by_status


```
Running 10s test @ http://localhost:8080
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   187.45ms   51.29ms 368.92ms   65.57%
    Req/Sec    11.22      4.32    20.00     71.57%
  424 requests in 10.03s, 612.89MB read
Requests/sec:     42.28
Transfer/sec:     61.12MB
```

---

## 04_read_all


```
Running 10s test @ http://localhost:8080
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   192.81ms   68.68ms 577.00ms   76.72%
    Req/Sec    11.27      4.82    30.00     68.49%
  416 requests in 10.09s, 602.78MB read
Requests/sec:     41.22
Transfer/sec:     59.72MB
```

---

## 05_healthz


```
Running 10s test @ http://localhost:8080
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   710.52us  742.49us  27.02ms   92.47%
    Req/Sec     3.01k   542.30     3.95k    74.50%
  120053 requests in 10.01s, 8.59MB read
Requests/sec:  11994.78
Transfer/sec:      0.86MB
```

---

