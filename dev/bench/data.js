window.BENCHMARK_DATA = {
  "lastUpdate": 1783107102817,
  "repoUrl": "https://github.com/dvprokofiev/arrangio-core",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "d@dvprokofiev.ru",
            "name": "dvprokofiev",
            "username": "dvprokofiev"
          },
          "committer": {
            "email": "d@dvprokofiev.ru",
            "name": "dvprokofiev",
            "username": "dvprokofiev"
          },
          "distinct": false,
          "id": "665ce12d9fe7ec0a1699836bd6107610596c6f89",
          "message": "fix(ci): create GH Pages page to show off benchmark results",
          "timestamp": "2026-06-10T13:49:15+03:00",
          "tree_id": "94c71c64f751f1a079b12a80d08c9b61f601bc1a",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/665ce12d9fe7ec0a1699836bd6107610596c6f89"
        },
        "date": 1781088620156,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "56175390 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.24,
            "unit": "ns/op",
            "extra": "56175390 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "56175390 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "56175390 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 810.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1487535 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 810.4,
            "unit": "ns/op",
            "extra": "1487535 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1487535 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1487535 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2620,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "458140 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2620,
            "unit": "ns/op",
            "extra": "458140 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "458140 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "458140 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 907.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1320357 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 907.8,
            "unit": "ns/op",
            "extra": "1320357 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1320357 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1320357 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "d@dvprokofiev.ru",
            "name": "dvprokofiev",
            "username": "dvprokofiev"
          },
          "committer": {
            "email": "d@dvprokofiev.ru",
            "name": "dvprokofiev",
            "username": "dvprokofiev"
          },
          "distinct": true,
          "id": "9dc11ec9a2bf4c757fc14dfd94b3c3e04524b21a",
          "message": "feat(rules): introduce `Rule` interface\n\n- `Type` - hard constraint (= -inf) and soft constraint\n- `Weight` - from 0 to 1, adds an ability to range rules from most important to least important\n- `Evaluate` - return value of rule violation, where 0 means rule is fully satisfied and 1 means that rule is not followed",
          "timestamp": "2026-07-03T22:30:34+03:00",
          "tree_id": "255afc4df7b64d4fbf076fbc4f5549a9831d7b70",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/9dc11ec9a2bf4c757fc14dfd94b3c3e04524b21a"
        },
        "date": 1783107102070,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 22.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52011086 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 22.56,
            "unit": "ns/op",
            "extra": "52011086 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52011086 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52011086 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 919.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1305772 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 919.3,
            "unit": "ns/op",
            "extra": "1305772 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1305772 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1305772 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2995,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "400400 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2995,
            "unit": "ns/op",
            "extra": "400400 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "400400 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "400400 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 942.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1270341 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 942.9,
            "unit": "ns/op",
            "extra": "1270341 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1270341 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1270341 times\n4 procs"
          }
        ]
      }
    ]
  }
}