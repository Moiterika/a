# a

<!-- coverage color = red ＜ 65% ≦ yellow ＜ 85%　≦ green -->

![Go](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic)
![Codecov_84.4%](https://img.shields.io/badge/coverage-84.4%-yellow)
[![License_MIT](https://img.shields.io/badge/license-MIT-blue)](https://github.com/Moiterika/a/blob/main/LICENSE)

## Abstruct

The `a` package includes simple Go utilities for collection operation. 
This package is inspired by the language integrated query (LINQ) in .NET Framework. 
There are no structs/interfaces equivalent to the `Enumerable` class or the `IEnumerable` interface, so deferred execution using the iterator pattern is not possible. 

### Currently implemented

- AggregateSelect
  - the original function
- AggregateSelectWithIndex
  - the original function
- All
- Any
- Chunk
- ChunkWithErr
- Contains
- Distinct
- DistinctBy
- DistinctByWithErr
- Filters
  - the original function
  - Filters() returns results corresponding to filters and the unfiltered result.
- FirstOrDefault
- GroupBy
- GroupByWithErr
- LastOrDefault
- MaxBy
- MaxByWithErr
- MinBy
- MinByWithErr
- OrderBy
- OrderByWithErr
- OrderByDescending
- OrderByDescendingWithErr
- Select
- SelectWithErr
- SelectMany
- Sort
  - support for multi-sorting with combinations of `OrderBy`, `ThenBy`, `OrderByDescending`, and `ThenByDescending`.
- Sum
- SumWithErr
- Switch
  - the original function
  - Switch() returns results corresponding to cases and the unfiltered result.
- ToMap
  - as an alternative to `ToDictionary`
- ToMapWithErr
- Where
- WhereSelect
  - the original function
  - WhereSelect() is convenient when using Where() and Select() at the same time.
- WhereSelectWithErr

### Todo

- Doc
- Implement
  - Aggregate
  - AsEnumerable
  - Average
  - Cast
  - Concat
  - Count
  - DefaultIfEmpty
  - ElementAt
  - ElementAtOrDefault
  - Empty
  - Except
  - ExceptBy
  - First
  - GroupJoin
  - Intersect
  - IntersectBy
  - Join
  - Last
  - LongCount
  - Max
  - Min
  - OfType
  - Order
  - OrderDescending
  - Prepend
  - Range
  - Repeat
  - Reverse
  - SequenceEqual
  - Single
  - SingleOrDefault
  - Skip
  - SkipLast
  - SkipWhile
  - Sum
    - for only number slice
  - Take
  - TakeLast
  - TakeWhile
  - ThenBy
  - ThenByDescending
  - ToArray
  - ToHashSet
  - ToList
  - ToLookup
  - TryGetNonEnumeratedCount
  - Union
  - UnionBy
  - Zip
- Test
  - ToMap
  - WhereSelect

### Tests

- How to measure code coverage
  - `go test -cover ./...`
- How to find not covered lines of code
  - `go test -cover ./... -coverprofile=cover.out`
  - `go tool cover -html=cover.out -o cover.html`
  - see `cover.html`
