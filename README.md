# lang

[![Go Version][version-img]][doc] [![GoDoc][doc-img]][doc] [![Build][ci-img]][ci] [![GoReport][report-img]][report]

**Package `lang` provides useful generic one-liners to work with variables, slices and maps**

```
go get -u github.com/maxbolgarin/lang
```

## Overview

`lang` is a lightweight utility library that adds a functional programming style to Go. It leverages Go's generics to provide type-safe operations for collections. This library saves you from writing boilerplate code for common operations, making your code more readable and maintainable.

## Features

### Slice Operations

- **Transformation**: `Map`, `Convert`, `ConvertWithErr`
- **Filtering**: `Filter`, `WithoutEmpty`, `NotEmpty`
- **Aggregation**: `Reduce`
- **Search**: `FindFirst`, `Contains`, `ContainsFunc`, `IndexOf`, `LastIndexOf`
- **Manipulation**: `Copy`, `SplitByChunkSize`, `Chunk`, `Flatten`, `Reverse`, `Distinct`, `Take`, `Skip`
- **Set Operations**: `Intersect`, `Union`, `Difference`
- **Iteration**: `ForEach`
- **Testing**: `All`, `Any`
- **Advanced**: `Partition`, `Compact`

### Map Operations

- **Creation**: `SliceToMap`, `SliceToMapByKey`, `Mapping`, `ConvertToMap`, `PairsToMap`, `ZipToMap`
- **Transformation**: `ConvertMap`, `ConvertMapWithErr`, `ConvertFromMap`, `ConvertFromMapWithErr`
- **Filtering**: `FilterMap`, `WithoutEmptyKeys`, `WithoutEmptyValues`, `NotEmptyMap`
- **Retrieval**: `Keys`, `KeysIf`, `Values`, `ValuesIf`
- **Manipulation**: `CopyMap`, `MergeMap`
- **Grouping**: `GroupBy`

## Examples

Here are a few examples to get you started:

### Filter even numbers
```go
numbers := []int{1, 2, 3, 4, 5}
evens := lang.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// evens = [2, 4]
```

### Transform a slice
```go
numbers := []int{1, 2, 3}
squares := lang.Convert(numbers, func(n int) int {
    return n * n
})
// squares = [1, 4, 9]
```

### Group items by a key
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 25},
}

byAge := lang.GroupBy(people, func(p Person) int {
    return p.Age
})
// byAge = map[int][]Person{
//     25: [{"Alice", 25}, {"Charlie", 25}],
//     30: [{"Bob", 30}],
// }
```

### Create a map from a slice
```go
users := []string{"user1", "user2", "user3"}
userMap := lang.SliceToMap(users, func(user string) (string, int) {
    return user, len(user)
})
// userMap = {"user1": 5, "user2": 5, "user3": 5}
```

### Merge multiple maps
```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}
merged := lang.MergeMap(map1, map2)
// merged = {"a": 1, "b": 3, "c": 4}
```

## Comparison with Native Go

Without generics, operations like filtering or mapping require writing full implementations each time:

```go
// Without lang
evens := []int{}
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}

// With lang
evens := lang.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
```

Add to your Go program a spice of functional languages: [docs][doc].

You also should try this package: [lo](https://github.com/samber/lo).


[version-img]: https://img.shields.io/badge/Go-%3E%3D%201.19-%23007d9c
[doc-img]: https://pkg.go.dev/badge/github.com/maxbolgarin/lang
[doc]: https://pkg.go.dev/github.com/maxbolgarin/lang
[ci-img]: https://github.com/maxbolgarin/lang/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/maxbolgarin/lang/actions
[report-img]: https://goreportcard.com/badge/github.com/maxbolgarin/lang
[report]: https://goreportcard.com/report/github.com/maxbolgarin/lang
