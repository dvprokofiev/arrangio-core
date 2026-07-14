window.BENCHMARK_DATA = {
  "lastUpdate": 1784043822748,
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
          "id": "208862b496c457c093e9eed11f19912ab39d5640",
          "message": "feat(rules): introduce `ProximityRule` interface\n\n- rule calculates attraction/distance scores to the closest object matching `Selector` scaled from 0.0 to 1.0",
          "timestamp": "2026-07-10T19:49:57+03:00",
          "tree_id": "7d2b524cd82346083e09772c7f88b4404ce77a26",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/208862b496c457c093e9eed11f19912ab39d5640"
        },
        "date": 1783702245010,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 22.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "53786943 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 22.57,
            "unit": "ns/op",
            "extra": "53786943 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "53786943 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "53786943 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 919.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1303534 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 919.6,
            "unit": "ns/op",
            "extra": "1303534 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1303534 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1303534 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2978,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "401698 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2978,
            "unit": "ns/op",
            "extra": "401698 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "401698 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "401698 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 945.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1269733 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 945.2,
            "unit": "ns/op",
            "extra": "1269733 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1269733 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1269733 times\n4 procs"
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
          "id": "33324f88f4d20be06b589005f7d465c4c8a0d20f",
          "message": "feat(rules): introduce `AxisRestrictionRule` interface\n\n- evaluates objects's placement boundaries along X, Y or Z spatial axis\n- supports relational operations\n- provides non-linear gradient on non-matched objects to guide mutation",
          "timestamp": "2026-07-10T20:53:43+03:00",
          "tree_id": "d49312f36e4708f06894836b98ff899a7a2130db",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/33324f88f4d20be06b589005f7d465c4c8a0d20f"
        },
        "date": 1783706076636,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 22.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51997591 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 22.99,
            "unit": "ns/op",
            "extra": "51997591 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51997591 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51997591 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 802.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1487860 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 802.1,
            "unit": "ns/op",
            "extra": "1487860 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1487860 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1487860 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2602,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "459658 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2602,
            "unit": "ns/op",
            "extra": "459658 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "459658 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "459658 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 908.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1318641 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 908.4,
            "unit": "ns/op",
            "extra": "1318641 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1318641 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1318641 times\n4 procs"
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
          "id": "fb3e5adc2e39f7da93a04f6e7b6b592bb77db57f",
          "message": "fix(rules): logical errors in axis restriction rule\n\n- use integers (0, 1, 2) in `Axis` instead of characters 'X', 'Y', 'Z'\n- if `OpNot` is matched, return 0.5 because we cannot provide further direction for this kind of operator\n- `if diff <= 0` is now out of case for `OpGt`, `OpGe`",
          "timestamp": "2026-07-11T21:33:52+03:00",
          "tree_id": "5632e99db068646a2f71e2f0eb463d00f1fbef48",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/fb3e5adc2e39f7da93a04f6e7b6b592bb77db57f"
        },
        "date": 1783794885240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.23,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54678554 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.23,
            "unit": "ns/op",
            "extra": "54678554 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54678554 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54678554 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 919.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1299825 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 919.1,
            "unit": "ns/op",
            "extra": "1299825 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1299825 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1299825 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 3005,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "398862 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 3005,
            "unit": "ns/op",
            "extra": "398862 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "398862 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "398862 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 1009,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 1009,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
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
          "id": "bd969a3bd72682b111bfa90fac751e560ea7b690",
          "message": "feat(rules): introduce `AlignmentRule` interface\n\n- align `subject` to nearby objects that match the `Selector`",
          "timestamp": "2026-07-11T22:03:05+03:00",
          "tree_id": "a8b24893bd341e5c04438106380877c38195dd62",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/bd969a3bd72682b111bfa90fac751e560ea7b690"
        },
        "date": 1783796643102,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.25,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "53243289 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.25,
            "unit": "ns/op",
            "extra": "53243289 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "53243289 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "53243289 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 919,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1232212 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 919,
            "unit": "ns/op",
            "extra": "1232212 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1232212 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1232212 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2998,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "402205 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2998,
            "unit": "ns/op",
            "extra": "402205 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "402205 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "402205 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 973.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1232192 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 973.6,
            "unit": "ns/op",
            "extra": "1232192 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1232192 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1232192 times\n4 procs"
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
          "id": "f64146222b585ce7e2bacf8769750e46048dbacb",
          "message": "fix(rules): use `uint8` instead of `int64` to store axes",
          "timestamp": "2026-07-14T16:34:33+03:00",
          "tree_id": "4134e61e459386d41feb25fa39d92a2dacea5e59",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/f64146222b585ce7e2bacf8769750e46048dbacb"
        },
        "date": 1784036145015,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "55625145 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.24,
            "unit": "ns/op",
            "extra": "55625145 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "55625145 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "55625145 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 917.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1297387 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 917.3,
            "unit": "ns/op",
            "extra": "1297387 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1297387 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1297387 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2983,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "401173 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2983,
            "unit": "ns/op",
            "extra": "401173 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "401173 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "401173 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 978.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1230054 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 978.7,
            "unit": "ns/op",
            "extra": "1230054 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1230054 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1230054 times\n4 procs"
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
          "id": "1ed218c815375cd7680cb4393eb4b398271cc69b",
          "message": "fix(rules): math error in axis restriction rule\n\n- if `val >= Ref` for 'less' operator, distance to valid zone (which is <= Ref - 1) is equal `val - (Ref - 1)`\n- if `val <= Ref` for 'greater' operator, distance to valid zone (which is >= Ref + 1) is equal `(Ref + 1) - val`",
          "timestamp": "2026-07-14T17:38:49+03:00",
          "tree_id": "0003cf2e904d345a979a2ef8fcdc5a31ce8ed936",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/1ed218c815375cd7680cb4393eb4b398271cc69b"
        },
        "date": 1784039980401,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 11.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "97555978 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 11.01,
            "unit": "ns/op",
            "extra": "97555978 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "97555978 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "97555978 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 531.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2280290 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 531.7,
            "unit": "ns/op",
            "extra": "2280290 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2280290 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2280290 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 1302,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "921859 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 1302,
            "unit": "ns/op",
            "extra": "921859 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "921859 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "921859 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 515.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2327985 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 515.6,
            "unit": "ns/op",
            "extra": "2327985 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2327985 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2327985 times\n4 procs"
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
          "id": "0f5f89c2068770cb7997c9f01154e24651ae1236",
          "message": "tests(rules): introduce 'framework' for rules tests\n\n- describe entities inserted to the grid by using `TestEntity` structure\n- `buildTestEntity` builds `Entity` from `TestEntity` with given `ID`, `Anchor`, `Tags` and size\n- `RunRuleTest` creates grid 2000x2000x2000, places neighbors from `RuleTestCase` and evaluates given rule",
          "timestamp": "2026-07-14T18:19:51+03:00",
          "tree_id": "407ec887596c47a3a5953ed30e0c971aed5ead1d",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/0f5f89c2068770cb7997c9f01154e24651ae1236"
        },
        "date": 1784042464276,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 22.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51370024 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 22.57,
            "unit": "ns/op",
            "extra": "51370024 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51370024 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51370024 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 923.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1306970 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 923.6,
            "unit": "ns/op",
            "extra": "1306970 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1306970 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1306970 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 3019,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "395938 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 3019,
            "unit": "ns/op",
            "extra": "395938 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "395938 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "395938 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 945.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1262497 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 945.8,
            "unit": "ns/op",
            "extra": "1262497 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1262497 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1262497 times\n4 procs"
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
          "id": "0f5f89c2068770cb7997c9f01154e24651ae1236",
          "message": "tests(rules): introduce 'framework' for rules tests\n\n- describe entities inserted to the grid by using `TestEntity` structure\n- `buildTestEntity` builds `Entity` from `TestEntity` with given `ID`, `Anchor`, `Tags` and size\n- `RunRuleTest` creates grid 2000x2000x2000, places neighbors from `RuleTestCase` and evaluates given rule",
          "timestamp": "2026-07-14T18:19:51+03:00",
          "tree_id": "407ec887596c47a3a5953ed30e0c971aed5ead1d",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/0f5f89c2068770cb7997c9f01154e24651ae1236"
        },
        "date": 1784042707574,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54525582 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.26,
            "unit": "ns/op",
            "extra": "54525582 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54525582 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54525582 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 917.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1300906 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 917.5,
            "unit": "ns/op",
            "extra": "1300906 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1300906 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1300906 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2985,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "402529 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2985,
            "unit": "ns/op",
            "extra": "402529 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "402529 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "402529 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 943.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1269679 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 943.7,
            "unit": "ns/op",
            "extra": "1269679 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1269679 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1269679 times\n4 procs"
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
          "id": "1ed218c815375cd7680cb4393eb4b398271cc69b",
          "message": "fix(rules): math error in axis restriction rule\n\n- if `val >= Ref` for 'less' operator, distance to valid zone (which is <= Ref - 1) is equal `val - (Ref - 1)`\n- if `val <= Ref` for 'greater' operator, distance to valid zone (which is >= Ref + 1) is equal `(Ref + 1) - val`",
          "timestamp": "2026-07-14T17:38:49+03:00",
          "tree_id": "0003cf2e904d345a979a2ef8fcdc5a31ce8ed936",
          "url": "https://github.com/dvprokofiev/arrangio-core/commit/1ed218c815375cd7680cb4393eb4b398271cc69b"
        },
        "date": 1784043821493,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGridInsert",
            "value": 21.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "55005207 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - ns/op",
            "value": 21.75,
            "unit": "ns/op",
            "extra": "55005207 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "55005207 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "55005207 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery",
            "value": 920.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1247346 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - ns/op",
            "value": 920.7,
            "unit": "ns/op",
            "extra": "1247346 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1247346 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1247346 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant",
            "value": 2999,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "401478 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - ns/op",
            "value": 2999,
            "unit": "ns/op",
            "extra": "401478 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "401478 times\n4 procs"
          },
          {
            "name": "BenchmarkGridInsert_Giant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "401478 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense",
            "value": 948.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1257162 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - ns/op",
            "value": 948.7,
            "unit": "ns/op",
            "extra": "1257162 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1257162 times\n4 procs"
          },
          {
            "name": "BenchmarkGridQuery_Dense - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1257162 times\n4 procs"
          }
        ]
      }
    ]
  }
}