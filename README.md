# a

<!-- coverage color = red ＜ 65% ≦ yellow ＜ 85%　≦ green -->

![Go](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic)
![Codecov](https://img.shields.io/badge/coverage-82.8%-yellow)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/Moiterika/a/blob/main/LICENSE)

## Abstruct

The `a` package is simple Go utilities for collection operation. 
This is inspired by the language integrated query (LINQ) in .NET Framework. 
This does not contain a struct equivalent to the `IEnumerable` interface, so deferred execution using the iterator pattern is not possible. 

### Currently implemented

- Any
- Chunk
- Contains
- Distinct
- Filters
  - the original function
  - Filters() returns results corresponding to filters and the unfiltered result.
- FirstOrDefault
- GroupBy
- LastOrDefault
- OrderBy
- OrderByDescending
- Select
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
  - All
  - AsEnumerable
  - Average
  - Cast
  - Concat
  - Count
  - DefaultIfEmpty
  - DistinctBy
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
  - MaxBy
  - Min
  - MinBy
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
  - Sum
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
  - Any
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
