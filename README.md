# fltcmp

[![Go Report Card](https://goreportcard.com/badge/github.com/erinpentecost/fltcmp)](https://goreportcard.com/report/github.com/erinpentecost/fltcmp)
[![Travis CI](https://travis-ci.org/erinpentecost/fltcmp.svg?branch=master)](https://travis-ci.org/erinpentecost/fltcmp.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/erinpentecost/fltcmp?status.svg)](https://godoc.org/github.com/erinpentecost/fltcmp)

Go float comparison using Units in the Last Place (ULPs). You can find a good write up of the method [here](https://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/).

When you test for float equality, you typically aren't looking for an *exact* match. What you might end up doing is testing for equality with some constant error threshold. The problem with that is the behaviour is way different when comparing very small floats and very large floats. ULP testing is an alternative.

## Benchmark

It's slow.

```sh
BenchmarkAlmostEqualDifferent-8         300000000                4.11 ns/op
BenchmarkAlmostEqualSame-8              300000000                3.98 ns/op
BenchmarkEpsilonSame-8                  2000000000               0.30 ns/op
BenchmarkEpsilonDifferent-8             2000000000               0.30 ns/op
```