# groot-bench

`groot-bench` gathers a few programs to benchmark the read/write performances of [groot](https://go-hep.org/x/hep/groot) _wrt_ ROOT/C++.


## input data

- `root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root` (1.8Gb, 23571931 entries)
- toy-data (`float64` and/or `[]float64`)

## results

### results - November 2024

- Go-HEP `v0.36` (Go-1.24@673a539170), `ROOT-6.32/06` (`g++-14.2.1`)

```
benchstat ./testdata/log-groot-v0.36.txt
goos: linux
goarch: amd64
pkg: github.com/go-hep/groot-bench
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
                                       │ ./testdata/log-groot-v0.36.txt │
                                       │             sec/op             │
ReadScalar/GoHEP/None-8                                     709.4m ± 2%
ReadScalar/ROOT-TreeBranch/None-8                            1.053 ± 2%
ReadScalar/ROOT-TreeReader/None-8                            1.266 ± 1%

ReadScalar/GoHEP/LZ4-8                                      697.8m ± 3%
ReadScalar/ROOT-TreeBranch/LZ4-8                             1.059 ± 2%
ReadScalar/ROOT-TreeReader/LZ4-8                             1.266 ± 1%

ReadScalar/GoHEP/Zlib-Lvl1-8                                728.8m ± 5%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                       1.063 ± 1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                       1.285 ± 1%

ReadScalar/GoHEP/Zlib-Lvl6-8                                709.3m ± 2%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                       1.060 ± 2%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                       1.287 ± 1%

ReadScalar/GoHEP/Zlib-Lvl9-8                                 1.349 ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                       2.200 ± 1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                       2.421 ± 1%

ReadSlices/GoHEP/None-8                                     244.3m ± 2%
ReadSlices/ROOT-TreeBranch/None-8                           435.0m ± 1%
ReadSlices/ROOT-TreeReader/None-8                           500.5m ± 1%

ReadSlices/GoHEP/LZ4-8                                      241.2m ± 1%
ReadSlices/ROOT-TreeBranch/LZ4-8                            439.9m ± 1%
ReadSlices/ROOT-TreeReader/LZ4-8                            506.7m ± 1%

ReadSlices/GoHEP/Zlib-Lvl1-8                                240.6m ± 1%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                      445.9m ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                      512.0m ± 1%

ReadSlices/GoHEP/Zlib-Lvl6-8                                242.8m ± 2%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                      446.2m ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                      526.9m ± 4%

ReadSlices/GoHEP/Zlib-Lvl9-8                                712.2m ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                       1.141 ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                       1.249 ± 1%

ReadCMSAll/GoHEP/Zlib-8                                      13.94 ± 2%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                            31.07 ± 2%

ReadCMSScalar/GoHEP/Zlib-8                                   3.059 ± 1%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                         8.236 ± 2%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                         7.124 ± 2%
geomean                                                      1.043

                                       │ ./testdata/log-groot-v0.36.txt │
                                       │             rss/op             │
ReadScalar/GoHEP/None-8                                    12.46M ±  1%
ReadScalar/ROOT-TreeBranch/None-8                          296.4M ±  0%
ReadScalar/ROOT-TreeReader/None-8                          296.5M ±  0%

ReadScalar/GoHEP/LZ4-8                                     12.41M ±  1%
ReadScalar/ROOT-TreeBranch/LZ4-8                           296.6M ±  0%
ReadScalar/ROOT-TreeReader/LZ4-8                           296.7M ±  0%

ReadScalar/GoHEP/Zlib-Lvl1-8                               12.49M ±  1%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                     296.6M ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                     296.7M ±  0%

ReadScalar/GoHEP/Zlib-Lvl6-8                               13.17M ± 14%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                     296.5M ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                     296.8M ±  0%

ReadScalar/GoHEP/Zlib-Lvl9-8                               13.88M ± 10%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                     296.6M ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                     296.7M ±  0%

ReadSlices/GoHEP/None-8                                    11.93M ±  0%
ReadSlices/ROOT-TreeBranch/None-8                          296.5M ±  0%
ReadSlices/ROOT-TreeReader/None-8                          296.7M ±  0%

ReadSlices/GoHEP/LZ4-8                                     12.45M ±  1%
ReadSlices/ROOT-TreeBranch/LZ4-8                           296.7M ±  0%
ReadSlices/ROOT-TreeReader/LZ4-8                           296.9M ±  0%

ReadSlices/GoHEP/Zlib-Lvl1-8                               15.01M ±  4%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                     355.1M ±  0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                     296.9M ±  0%

ReadSlices/GoHEP/Zlib-Lvl6-8                               15.02M ± 10%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                     355.2M ±  0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                     355.5M ±  0%

ReadSlices/GoHEP/Zlib-Lvl9-8                               14.42M ±  7%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                     296.5M ±  0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                     296.8M ±  0%

ReadCMSAll/GoHEP/Zlib-8                                    114.4M ±  0%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                          282.0M ±  0%

ReadCMSScalar/GoHEP/Zlib-8                                 27.23M ±  6%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                       253.3M ±  0%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                       253.6M ±  0%
geomean                                                    111.6M

                                       │ ./testdata/log-groot-v0.36.txt │
                                       │            vmem/op             │
ReadScalar/GoHEP/None-8                                     1.232G ± 0%
ReadScalar/ROOT-TreeBranch/None-8                           452.0M ± 0%
ReadScalar/ROOT-TreeReader/None-8                           451.9M ± 0%

ReadScalar/GoHEP/LZ4-8                                      1.232G ± 0%
ReadScalar/ROOT-TreeBranch/LZ4-8                            452.0M ± 0%
ReadScalar/ROOT-TreeReader/LZ4-8                            451.9M ± 0%

ReadScalar/GoHEP/Zlib-Lvl1-8                                1.233G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                      452.0M ± 0%

ReadScalar/GoHEP/Zlib-Lvl6-8                                1.233G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                      452.0M ± 0%

ReadScalar/GoHEP/Zlib-Lvl9-8                                1.233G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                      452.0M ± 0%

ReadSlices/GoHEP/None-8                                     1.232G ± 0%
ReadSlices/ROOT-TreeBranch/None-8                           452.3M ± 0%
ReadSlices/ROOT-TreeReader/None-8                           452.3M ± 0%

ReadSlices/GoHEP/LZ4-8                                      1.233G ± 0%
ReadSlices/ROOT-TreeBranch/LZ4-8                            452.3M ± 0%
ReadSlices/ROOT-TreeReader/LZ4-8                            452.3M ± 0%

ReadSlices/GoHEP/Zlib-Lvl1-8                                1.233G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                      510.9M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                      452.3M ± 0%

ReadSlices/GoHEP/Zlib-Lvl6-8                                1.233G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                      510.9M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                      511.0M ± 0%

ReadSlices/GoHEP/Zlib-Lvl9-8                                1.233G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                      452.3M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                      452.3M ± 0%

ReadCMSAll/GoHEP/Zlib-8                                     1.299G ± 0%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                           474.2M ± 0%

ReadCMSScalar/GoHEP/Zlib-8                                  1.233G ± 0%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                        414.3M ± 0%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                        414.4M ± 0%
geomean                                                     643.0M

                                       │ ./testdata/log-groot-v0.36.txt │
                                       │            wall/op             │
ReadScalar/GoHEP/None-8                                     721.1m ± 3%
ReadScalar/ROOT-TreeBranch/None-8                            1.053 ± 2%
ReadScalar/ROOT-TreeReader/None-8                            1.265 ± 1%

ReadScalar/GoHEP/LZ4-8                                      693.7m ± 2%
ReadScalar/ROOT-TreeBranch/LZ4-8                             1.058 ± 2%
ReadScalar/ROOT-TreeReader/LZ4-8                             1.265 ± 1%

ReadScalar/GoHEP/Zlib-Lvl1-8                                725.8m ± 5%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                       1.062 ± 1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                       1.284 ± 1%

ReadScalar/GoHEP/Zlib-Lvl6-8                                709.6m ± 2%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                       1.060 ± 2%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                       1.287 ± 1%

ReadScalar/GoHEP/Zlib-Lvl9-8                                 1.349 ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                       2.200 ± 1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                       2.421 ± 1%

ReadSlices/GoHEP/None-8                                     241.9m ± 2%
ReadSlices/ROOT-TreeBranch/None-8                           435.2m ± 2%
ReadSlices/ROOT-TreeReader/None-8                           499.4m ± 2%

ReadSlices/GoHEP/LZ4-8                                      239.4m ± 3%
ReadSlices/ROOT-TreeBranch/LZ4-8                            436.4m ± 2%
ReadSlices/ROOT-TreeReader/LZ4-8                            504.9m ± 1%

ReadSlices/GoHEP/Zlib-Lvl1-8                                238.0m ± 2%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                      442.2m ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                      512.9m ± 3%

ReadSlices/GoHEP/Zlib-Lvl6-8                                241.9m ± 1%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                      444.1m ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                      527.4m ± 4%

ReadSlices/GoHEP/Zlib-Lvl9-8                                711.2m ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                       1.141 ± 1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                       1.249 ± 1%

ReadCMSAll/GoHEP/Zlib-8                                      13.94 ± 2%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                            31.07 ± 2%

ReadCMSScalar/GoHEP/Zlib-8                                   3.059 ± 1%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                         8.236 ± 2%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                         7.123 ± 2%
geomean                                                      1.041
```

- Go-HEP `v0.35` (Go-1.24@673a539170), `ROOT-6.32/06` (`g++-14.2.1`)

```
$> benchstat ./testdata/log-groot-v0.35.txt
goos: linux
goarch: amd64
pkg: github.com/go-hep/groot-bench
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
                                       │ ./testdata/log-groot-v0.35.txt │
                                       │             sec/op             │
ReadScalar/GoHEP/None-8                                    649.0m ±  2%
ReadScalar/ROOT-TreeBranch/None-8                           1.043 ±  1%
ReadScalar/ROOT-TreeReader/None-8                           1.258 ±  1%

ReadScalar/GoHEP/LZ4-8                                     656.4m ±  1%
ReadScalar/ROOT-TreeBranch/LZ4-8                            1.049 ±  1%
ReadScalar/ROOT-TreeReader/LZ4-8                            1.260 ±  1%

ReadScalar/GoHEP/Zlib-Lvl1-8                               648.6m ±  3%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                      1.051 ±  1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                      1.269 ±  1%

ReadScalar/GoHEP/Zlib-Lvl6-8                               658.8m ±  2%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                      1.050 ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                      1.260 ±  1%

ReadScalar/GoHEP/Zlib-Lvl9-8                                1.338 ±  0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                      2.186 ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                      2.417 ±  1%

ReadSlices/GoHEP/None-8                                    220.7m ±  1%
ReadSlices/ROOT-TreeBranch/None-8                          429.8m ±  1%
ReadSlices/ROOT-TreeReader/None-8                          496.3m ±  1%

ReadSlices/GoHEP/LZ4-8                                     220.2m ±  1%
ReadSlices/ROOT-TreeBranch/LZ4-8                           434.6m ±  0%
ReadSlices/ROOT-TreeReader/LZ4-8                           591.0m ±  2%

ReadSlices/GoHEP/Zlib-Lvl1-8                               219.4m ±  2%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                     440.6m ±  1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                     517.9m ± 16%

ReadSlices/GoHEP/Zlib-Lvl6-8                               224.3m ±  2%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                     444.5m ±  0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                     507.8m ±  1%

ReadSlices/GoHEP/Zlib-Lvl9-8                               710.8m ±  0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                      1.147 ±  1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                      1.220 ±  1%

ReadCMSAll/GoHEP/Zlib-8                                     13.14 ±  1%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                           30.19 ±  1%

ReadCMSScalar/GoHEP/Zlib-8                                  2.928 ±  1%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                        7.968 ±  1%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                        6.950 ±  1%
geomean                                                     1.015

                                       │ ./testdata/log-groot-v0.35.txt │
                                       │             rss/op             │
ReadScalar/GoHEP/None-8                                     324.0M ± 0%
ReadScalar/ROOT-TreeBranch/None-8                           296.4M ± 0%
ReadScalar/ROOT-TreeReader/None-8                           296.4M ± 0%

ReadScalar/GoHEP/LZ4-8                                      323.6M ± 0%
ReadScalar/ROOT-TreeBranch/LZ4-8                            296.5M ± 0%
ReadScalar/ROOT-TreeReader/LZ4-8                            296.7M ± 0%

ReadScalar/GoHEP/Zlib-Lvl1-8                                323.6M ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                      296.6M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                      296.8M ± 0%

ReadScalar/GoHEP/Zlib-Lvl6-8                                324.1M ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                      296.5M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                      296.7M ± 0%

ReadScalar/GoHEP/Zlib-Lvl9-8                                307.3M ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                      296.5M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                      296.7M ± 0%

ReadSlices/GoHEP/None-8                                     201.3M ± 1%
ReadSlices/ROOT-TreeBranch/None-8                           296.5M ± 0%
ReadSlices/ROOT-TreeReader/None-8                           296.7M ± 0%

ReadSlices/GoHEP/LZ4-8                                      201.1M ± 1%
ReadSlices/ROOT-TreeBranch/LZ4-8                            296.6M ± 0%
ReadSlices/ROOT-TreeReader/LZ4-8                            296.8M ± 0%

ReadSlices/GoHEP/Zlib-Lvl1-8                                201.1M ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                      355.2M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                      296.8M ± 0%

ReadSlices/GoHEP/Zlib-Lvl6-8                                201.9M ± 1%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                      355.3M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                      355.5M ± 0%

ReadSlices/GoHEP/Zlib-Lvl9-8                                184.9M ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                      296.4M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                      296.8M ± 0%

ReadCMSAll/GoHEP/Zlib-8                                     1.751G ± 0%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                           281.9M ± 0%

ReadCMSScalar/GoHEP/Zlib-8                                  389.4M ± 0%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                        253.4M ± 0%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                        253.5M ± 0%
geomean                                                     301.6M

                                       │ ./testdata/log-groot-v0.35.txt │
                                       │            vmem/op             │
ReadScalar/GoHEP/None-8                                     1.546G ± 0%
ReadScalar/ROOT-TreeBranch/None-8                           452.0M ± 0%
ReadScalar/ROOT-TreeReader/None-8                           451.9M ± 0%

ReadScalar/GoHEP/LZ4-8                                      1.546G ± 0%
ReadScalar/ROOT-TreeBranch/LZ4-8                            452.0M ± 0%
ReadScalar/ROOT-TreeReader/LZ4-8                            452.0M ± 0%

ReadScalar/GoHEP/Zlib-Lvl1-8                                1.545G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                      452.0M ± 0%

ReadScalar/GoHEP/Zlib-Lvl6-8                                1.546G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                      452.0M ± 0%

ReadScalar/GoHEP/Zlib-Lvl9-8                                1.528G ± 0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                      452.0M ± 0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                      452.0M ± 0%

ReadSlices/GoHEP/None-8                                     1.424G ± 0%
ReadSlices/ROOT-TreeBranch/None-8                           452.3M ± 0%
ReadSlices/ROOT-TreeReader/None-8                           452.3M ± 0%

ReadSlices/GoHEP/LZ4-8                                      1.422G ± 0%
ReadSlices/ROOT-TreeBranch/LZ4-8                            452.3M ± 0%
ReadSlices/ROOT-TreeReader/LZ4-8                            452.3M ± 0%

ReadSlices/GoHEP/Zlib-Lvl1-8                                1.422G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                      510.9M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                      452.3M ± 0%

ReadSlices/GoHEP/Zlib-Lvl6-8                                1.422G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                      510.9M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                      511.0M ± 0%

ReadSlices/GoHEP/Zlib-Lvl9-8                                1.405G ± 0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                      452.3M ± 0%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                      452.3M ± 0%

ReadCMSAll/GoHEP/Zlib-8                                     3.075G ± 0%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                           474.2M ± 0%

ReadCMSScalar/GoHEP/Zlib-8                                  3.009G ± 0%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                        414.3M ± 0%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                        414.4M ± 0%
geomean                                                     712.2M

                                       │ ./testdata/log-groot-v0.35.txt │
                                       │            wall/op             │
ReadScalar/GoHEP/None-8                                    650.5m ±  3%
ReadScalar/ROOT-TreeBranch/None-8                           1.043 ±  1%
ReadScalar/ROOT-TreeReader/None-8                           1.257 ±  1%

ReadScalar/GoHEP/LZ4-8                                     655.1m ±  3%
ReadScalar/ROOT-TreeBranch/LZ4-8                            1.048 ±  1%
ReadScalar/ROOT-TreeReader/LZ4-8                            1.260 ±  1%

ReadScalar/GoHEP/Zlib-Lvl1-8                               654.3m ±  5%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl1-8                      1.051 ±  1%
ReadScalar/ROOT-TreeReader/Zlib-Lvl1-8                      1.269 ±  1%

ReadScalar/GoHEP/Zlib-Lvl6-8                               661.2m ±  3%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl6-8                      1.050 ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl6-8                      1.260 ±  1%

ReadScalar/GoHEP/Zlib-Lvl9-8                                1.337 ±  0%
ReadScalar/ROOT-TreeBranch/Zlib-Lvl9-8                      2.185 ±  0%
ReadScalar/ROOT-TreeReader/Zlib-Lvl9-8                      2.417 ±  1%

ReadSlices/GoHEP/None-8                                    223.2m ±  3%
ReadSlices/ROOT-TreeBranch/None-8                          426.0m ±  1%
ReadSlices/ROOT-TreeReader/None-8                          496.3m ±  2%

ReadSlices/GoHEP/LZ4-8                                     220.3m ±  5%
ReadSlices/ROOT-TreeBranch/LZ4-8                           434.3m ±  1%
ReadSlices/ROOT-TreeReader/LZ4-8                           596.9m ±  5%

ReadSlices/GoHEP/Zlib-Lvl1-8                               219.5m ±  5%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl1-8                     437.6m ±  1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl1-8                     512.3m ± 17%

ReadSlices/GoHEP/Zlib-Lvl6-8                               228.8m ±  3%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl6-8                     446.1m ±  1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl6-8                     507.1m ±  1%

ReadSlices/GoHEP/Zlib-Lvl9-8                               711.1m ±  0%
ReadSlices/ROOT-TreeBranch/Zlib-Lvl9-8                      1.147 ±  1%
ReadSlices/ROOT-TreeReader/Zlib-Lvl9-8                      1.220 ±  1%

ReadCMSAll/GoHEP/Zlib-8                                     13.13 ±  1%
ReadCMSAll/ROOT-TreeBranch/Zlib-8                           30.20 ±  1%

ReadCMSScalar/GoHEP/Zlib-8                                  2.928 ±  1%
ReadCMSScalar/ROOT-TreeBranch/Zlib-8                        7.967 ±  1%
ReadCMSScalar/ROOT-TreeReader/Zlib-8                        6.950 ±  1%
geomean                                                     1.016
```

### results - May 2020

- Go-HEP `v0.27.0` (Go-1.14), `ROOT-6.20/04` (g++-10.1.0) (both on local file)

```
name                                  time/op
ReadCMSScalar/GoHEP/Zlib-8            3.92s ± 2%  // only read scalar data
ReadCMSScalar/ROOT-TreeBranch/Zlib-8  7.98s ± 2%  // ditto
ReadCMSScalar/ROOT-TreeReader/Zlib-8  6.60s ± 2%  // ditto

name                                  time/op
ReadCMSAll/GoHEP/Zlib-8               18.4s ± 1%  // read all branches
ReadCMSAll/ROOT-TreeBranch/Zlib-8     30.4s ± 2%  // ditto
ReadCMSAll/ROOT-TreeReader/Zlib-8     [N/A]       // comparison meaningless (b/c of loading-on-demand)

name                               time/op
ReadScalar/GoHEP/None-8            705ms ± 2%
ReadScalar/ROOT-TreeBranch/None-8  941ms ± 2%
ReadScalar/ROOT-TreeReader/None-8  1.39s ± 2%
ReadScalar/GoHEP/LZ4-8             751ms ± 2%
ReadScalar/ROOT-TreeBranch/LZ4-8   1.09s ± 2%
ReadScalar/ROOT-TreeReader/LZ4-8   1.53s ± 1%
ReadScalar/GoHEP/Zlib-8            1.31s ± 1%
ReadScalar/ROOT-TreeBranch/Zlib-8  2.18s ± 1%
ReadScalar/ROOT-TreeReader/Zlib-8  2.62s ± 1%
```

- Go-HEP `v0.26.1-0.20200511085556-0f7b59f24c5e`

```
name                     time/op
ReadScalar/GoHEP/None-8  1.27s ± 3%
ReadScalar/GoHEP/LZ4-8   1.39s ± 1%
ReadScalar/GoHEP/Zlib-8  4.10s ± 1%
```

## toy-data

### generation

```
$> make binaries
$> ./bin/gen-data-scalar.go -h
Usage of gen-data-scalar:
  -cpu-profile string
    	path to the output CPU profile
  -lvl int
    	compression level to use (if any) (default -1)
  -nevts int
    	number of events to generate (default 10000000)
  -o string
    	path to output ROOT file to generate (default "scalar.root")
  -seed uint
    	seed for random number generation (default 1234)
  -t string
    	name of the output ROOT tree to generate (default "tree")
  -zip string
    	compression to use (if any)
```

#### no-compression

```
./bin/gen-data-scalar -zip=none -o ./testdata/scalar-none.root
```

#### lz4

```
./bin/gen-data-scalar -zip=lz4 -lvl=0 -o ./testdata/scalar-lz4-0.root
./bin/gen-data-scalar -zip=lz4 -lvl=1 -o ./testdata/scalar-lz4-1.root
./bin/gen-data-scalar -zip=lz4 -lvl=6 -o ./testdata/scalar-lz4-6.root
./bin/gen-data-scalar -zip=lz4 -lvl=9 -o ./testdata/scalar-lz4-9.root
```

#### zlib

```
./bin/gen-data-scalar -zip=zlib -lvl=0 -o ./testdata/scalar-zlib-0.root
./bin/gen-data-scalar -zip=zlib -lvl=1 -o ./testdata/scalar-zlib-1.root
./bin/gen-data-scalar -zip=zlib -lvl=2 -o ./testdata/scalar-zlib-2.root
./bin/gen-data-scalar -zip=zlib -lvl=3 -o ./testdata/scalar-zlib-3.root
./bin/gen-data-scalar -zip=zlib -lvl=6 -o ./testdata/scalar-zlib-6.root
./bin/gen-data-scalar -zip=zlib -lvl=9 -o ./testdata/scalar-zlib-9.root
```

#### zstd

```
./bin/gen-data-scalar -zip=zstd -o ./testdata/scalar-zstd.root
```
