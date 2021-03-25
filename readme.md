## 2021-3-25
```
GithubAPI Routes: 205
   cottonRouter:     95352 bytes
 BeegoMuxRouter:    108848 bytes
     BoneRouter:    102888 bytes
      ChiRouter:     71240 bytes
     HttpRouter:     36016 bytes
       trie-mux:    132456 bytes
      GoRouter1:     83320 bytes
goos: darwin
goarch: amd64
pkg: cottonbench
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkHttpRouterWithGithubAPI-8                 39334             31510 ns/op           13856 B/op        169 allocs/op
BenchmarkCottonRouterWithGithubAPI-8               34222             35289 ns/op               0 B/op          0 allocs/op
BenchmarkBeegoMuxRouterWithGithubAPI-8             10018            120020 ns/op          139056 B/op       1050 allocs/op
BenchmarkBoneRouterWithGithubAPI-8                   636           1861204 ns/op          744017 B/op       8893 allocs/op
BenchmarkTrieMuxRouterWithGithubAPI-8              17384             69513 ns/op           66624 B/op        543 allocs/op
BenchmarkGoRouter1WithGithubAPI-8                     61          18119023 ns/op        14473757 B/op     133388 allocs/op
PASS
ok      cottonbench     11.384s
```

## 2021-3-21
```
GithubAPI Routes: 205
   cottonRouter:     93384 bytes
 BeegoMuxRouter:    112400 bytes
     BoneRouter:    100712 bytes
      ChiRouter:     73520 bytes
     HttpRouter:     36064 bytes
       trie-mux:    132248 bytes
      GoRouter1:     83536 bytes
goos: darwin
goarch: amd64
pkg: cottonbench
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkHttpRouterWithGithubAPI-8                 66828             17910 ns/op               0 B/op          0 allocs/op
BenchmarkCottonRouterWithGithubAPI-8               35174             33773 ns/op               0 B/op          0 allocs/op
BenchmarkBeegoMuxRouterWithGithubAPI-8             10286            116937 ns/op          139056 B/op       1050 allocs/op
BenchmarkBoneRouterWithGithubAPI-8                   649           1819208 ns/op          744019 B/op       8893 allocs/op
BenchmarkTrieMuxRouterWithGithubAPI-8              18014             66955 ns/op           66624 B/op        543 allocs/op
BenchmarkGoRouter1WithGithubAPI-8                     63          18027981 ns/op        14413968 B/op     133154 allocs/op
PASS
ok      cottonbench     9.844s
```

## 2021-3-19
```
tonnydeMBP:cotton-bench tonny$ go test -bench=.
GithubAPI Routes: 203
GithubAPI2 Routes: 203
           cottonRouter:     93080 bytes
         BeegoMuxRouter:    110008 bytes
             BoneRouter:     99672 bytes
             HttpRouter:     35768 bytes
               trie-mux:    132552 bytes
              GoRouter1:     82808 bytes
              GoRouter2:     83304 bytes
goos: darwin
goarch: amd64
pkg: cottonbench
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkCottonRouterWithGithubAPI-8                1543            792932 ns/op  598146 B/op        8100 allocs/op
BenchmarkBeegoMuxRouterWithGithubAPI-8              8113            139834 ns/op  143952 B/op        1205 allocs/op
BenchmarkBoneRouterWithGithubAPI-8                   627           1799618 ns/op  729361 B/op        8787 allocs/op
BenchmarkTrieMuxRouterWithGithubAPI-8              13953             86143 ns/op   72384 B/op         704 allocs/op
BenchmarkHttpRouterWithGithubAPI-8                 25250             46397 ns/op   20320 B/op         334 allocs/op
BenchmarkGoRouter1WithGithubAPI-8                     61          17794924 ns/op14362332 B/op      132470 allocs/op
PASS
ok      cottonbench     10.259s
tonnydeMBP:cotton-bench tonny$ 
```

