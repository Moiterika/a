# a

<!-- coverage color = red ＜ 65% ≦ yellow ＜ 85%　≦ green -->

![Go](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic)
![Codecov](https://img.shields.io/badge/coverage-72.6%-yellow)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/Moiterika/a/blob/main/LICENSE)

## Abstruct

The `a` package incluedes simple Go utilities of Moiterika LLC. 
This package is inspired by the language integrated query (LINQ) in .NET Framework. 
This package does not contain a struct equivalent to the `IEnumerable` interface, so deferred execution using the iterator pattern is not possible. 

### Currently implemented

- Any
- Chunk
- Exists
- Filters
- FirstOrDefault
- GroupBy
- LastOrDefault
- Select
- ToMap
- Where
- WhereSelect

### Todo

- Test
  - Any
  - Exists
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
