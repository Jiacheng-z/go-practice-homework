

### SET/GET 10byte 

```
====== SET ======                                                   
  100000 requests completed in 1.63 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 61274.51 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.445     0.088     0.375     0.791     1.631    10.943

====== GET ======                                                   
  100000 requests completed in 1.60 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 62383.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.422     0.040     0.367     0.783     1.383     3.535
```

### SET/GET 20byte

```
====== SET ======                                                   
  100000 requests completed in 1.87 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 53533.19 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.509     0.104     0.383     1.119     2.175     9.191

====== GET ======                                                   
  100000 requests completed in 1.70 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 58719.91 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.448     0.152     0.375     0.919     1.695     4.767
```

### SET/GET 50byte

```
====== SET ======                                                   
  100000 requests completed in 1.72 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 58173.36 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.452     0.048     0.375     0.959     1.831     6.023

====== GET ======                                                   
  100000 requests completed in 1.80 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 55524.71 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.481     0.160     0.383     1.063     1.727     5.119
```

### SET/GET 100byte

```
====== SET ======                                                   
  100000 requests completed in 1.87 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 53619.30 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.501     0.168     0.383     1.135     2.079     6.791

====== GET ======                                                   
  100000 requests completed in 1.74 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 57372.34 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.471     0.080     0.375     1.031     1.783     6.591
```

### SET/GET 200byte

```
====== SET ======                                                   
  100000 requests completed in 1.75 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 57110.22 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.481     0.144     0.375     1.079     1.783     5.207

====== GET ======                                                   
  100000 requests completed in 1.98 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no```

Summary:
  throughput summary: 50505.05 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.546     0.096     0.383     1.391     2.303    48.351
```

### SET/GET 1000byte

```
====== SET ======                                                   
  100000 requests completed in 2.04 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 49140.05 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.567     0.168     0.423     1.279     2.207     8.687

====== GET ======                                                   
  100000 requests completed in 2.05 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 48875.86 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.562     0.168     0.415     1.311     2.407     7.311
```

### SET/GET 5000byte

```
====== SET ======                                                   
  100000 requests completed in 2.08 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 48076.93 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.557     0.104     0.439     1.271     2.119     5.863

====== GET ======                                                   
  100000 requests completed in 1.94 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 51466.80 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.534     0.072     0.375     1.111     1.959    91.455
```

```
1w 10byte  = 90.6304
used_memory:1990160
used_memory_human:1.90M

1w 20byte = 101.1072
used_memory:1083856
used_memory_human:1.03M
used_memory:2094928
used_memory_human:2.00M

1w 50byte = 133.1072
used_memory:1108544
used_memory_human:1.06M
used_memory:2439616
used_memory_human:2.33M

1w 100byte = 189.1072
used_memory:1108544
used_memory_human:1.06M
used_memory:2999616
used_memory_human:2.86M

1w 200byte = 301.1072
used_memory:1108544
used_memory_human:1.06M
used_memory:4119616
used_memory_human:3.93M

1w 1000byte = 1101.1072
used_memory:1108544
used_memory_human:1.06M
used_memory:12119616
used_memory_human:11.56M

1w 5000byte = 5196.6264
used_memory:1109248
used_memory_human:1.06M
used_memory:53075512
used_memory_human:50.62M
```
