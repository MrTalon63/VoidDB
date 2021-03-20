VoidDB
---
Basic in-memory key-value storage with an HTTP interface. Use at own risk.

### Endpoints
| Method | URI       | Redis equivalent | Description                                           |
|--------|-----------|------------------|-------------------------------------------------------|
| GET    | /         | N/A              | Welcome page                                          |
| GET    | /stats    | N/A              | App statistics in the Prometheus format               |
| GET    | /dump     | N/A              | All the data exported in the JSON format (key=base64) |
| POST   | /db/{key} | SET              | Set the value for the {key} to the body contents      |
| GET    | /db/{key} | GET              | Get the data for the specified {key}                  |
| DELETE | /db/{key} | DEL              | Delete the data by the specified {key}                |

### Notes
- The data is stored in the `save.json` file.
- The data is being saved only every five seconds and on program exit (SIGINT/SIGTERM).

### Stats
| Key                    | Comment                    |
|------------------------|----------------------------|
| voiddb_entries_total   | current data entries count |
| voiddb_runtime_seconds | runtime in seconds         |
| voiddb_stat_read       | reads from the startup     |
| voiddb_stat_write      | writes from the startup    |
| voiddb_stat_delete     | deletions from the startup |

### Benchmarks
Performed on Ryzen 5 3600 system using ApacheBench with 10 concurrent clients.
```console
# SET test (36 characters payload)
$ ab -p test.txt -n 1000000 -c 10 http://127.0.0.1:8080/db/test-key
Requests per second:    39467.91 [#/sec] (mean)
Time per request:       0.253 [ms] (mean)
Time per request:       0.025 [ms] (mean, across all concurrent requests)

# GET test (reading same 36 characters of data)
$ ab -n 1000000 -c 10 http://127.0.0.1:8080/db/test-key
Requests per second:    39656.23 [#/sec] (mean)
Time per request:       0.252 [ms] (mean)
Time per request:       0.025 [ms] (mean, across all concurrent requests)
```

### Features
- [x] Basic operations: SET/GET/DEL
- [x] Ability to download a snapshot of database (dump)
- [x] Auto-save feature (periodic and on an exit)
- [ ] Connection authentication (header)
- [ ] Config (auto-save frequency, server port, credentials)
