# a

<!-- coverage color = red ＜ 65% ≦ yellow ＜ 85%　≦ green -->

![Go](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic)
![Codecov](https://img.shields.io/badge/coverage-91.3%-green)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/Moiterika/a/blob/main/LICENSE)

## Abstruct

The `a` package is simple Go utilities for collection operation. 
This is inspired by the language integrated query (LINQ) in .NET Framework. 
This does not contain a struct equivalent to the `IEnumerable` interface, so deferred execution using the iterator pattern is not possible. 

### Currently implemented

- All
- Any
- Chunk
- Contains
- Distinct
- DistinctBy
- Filters
  - the original function
  - Filters() returns results corresponding to filters and the unfiltered result.
- FirstOrDefault
- GroupBy
- LastOrDefault
- MaxBy
- MinBy
- OrderBy
- OrderByDescending
- Select
- Sum
  - TODO
    - Sum for only number slice
- ToMap
  - as an alternative to `ToDictionary`
- Where
- WhereSelect
  - the original function
  - WhereSelect() is convenient when using Where() and Select() at the same time.

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
  - SelectMany
  - SequenceEqual
  - Single
  - SingleOrDefault
  - Skip
  - SkipLast
  - SkipWhile
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
  - GroupBy
  - ToMap
  - WhereSelect

### Tests

- How to measure code coverage
  - `go test -cover ./...`
- How to find not covered lines of code
  - `go test -cover ./... -coverprofile=cover.out`
  - `go tool cover -html=cover.out -o cover.html`
  - see `cover.html`
