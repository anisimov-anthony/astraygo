# WRK Benchmark Report - AstrayGo

**Generated:** 2026-01-29 09:25:47 UTC

## System Specifications

**CPU:**
Model name:                              AMD EPYC 9474F 48-Core Processor
- Cores: 24
- Threads: 24

**Memory:**
Total: 47Gi, Available: 44Gi

**Operating System:**
Ubuntu 22.04.5 LTS

**Kernel:**
5.15.0-164-generic

**Architecture:**
x86_64

---

## 01_warmup_post


```
Running 30s test @ http://localhost:8080
  24 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   339.40ms   25.28ms 446.97ms   90.47%
    Req/Sec     1.33k     1.05k   15.43k    66.45%
  878811 requests in 30.10s, 189.96MB read
Requests/sec:  29198.67
Transfer/sec:      6.31MB
```

---

## 02_read_by_id


```
Running 30s test @ http://localhost:8080
  24 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   152.71ms   57.69ms   1.12s    97.80%
    Req/Sec     2.85k   552.97    18.26k    79.19%
  2015508 requests in 30.09s, 426.08MB read
Requests/sec:  66971.91
Transfer/sec:     14.16MB
```

---

## 03_read_by_status


```
Running 30s test @ http://localhost:8080
  24 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.31s   126.67ms   2.00s    94.49%
    Req/Sec   285.14    244.07     2.38k    70.02%
  185353 requests in 30.10s, 17.11GB read
Requests/sec:   6157.90
Transfer/sec:    581.98MB
```

---

## 04_read_all


```
Running 30s test @ http://localhost:8080
  24 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.29s   137.75ms   2.00s    94.76%
    Req/Sec   339.45    366.91     2.94k    86.83%
  187500 requests in 30.10s, 17.31GB read
Requests/sec:   6229.66
Transfer/sec:    588.76MB
```

---

## 05_healthz


```
Running 30s test @ http://localhost:8080
  24 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    46.31ms   48.26ms 860.36ms   97.07%
    Req/Sec    10.11k     1.70k   21.50k    89.90%
  7235401 requests in 30.09s, 517.52MB read
Requests/sec: 240432.39
Transfer/sec:     17.20MB
```

---
