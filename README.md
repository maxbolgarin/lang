# lang

[![Go Version][version-img]][doc] [![GoDoc][doc-img]][doc] [![Build][ci-img]][ci] [![Coverage][coverage-img]][coverage] [![GoReport][report-img]][report] [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

**Package `lang` provides useful generic one-liners to work with variables, slices and maps**

```bash
go get -u github.com/maxbolgarin/lang
```

## Overview

`lang` is a comprehensive utility library that brings functional programming concepts to Go. It leverages Go's generics to provide type-safe operations for collections, error handling, and data transformations. This library eliminates boilerplate code and makes your Go code more expressive and maintainable.

## Table of Contents

- [Core Utilities](#core-utilities)
- [Slice Operations](#slice-operations)
- [Map Operations](#map-operations)
- [Error Handling](#error-handling)
- [Type Safety](#type-safety)
- [Real-World Examples](#real-world-examples)
- [Performance Notes](#performance-notes)
- [Best Practices](#best-practices)

## Core Utilities

### Pointer and Value Helpers

These functions help you work with pointers and provide safe defaults:

```go
// Working with pointers
name := lang.Ptr("John")           // *string pointing to "John"
age := lang.Ptr(25)                // *int pointing to 25

// Safe pointer dereferencing
var nilPtr *string
value := lang.Deref(nilPtr)        // "" (zero value)
value = lang.Deref(lang.Ptr("Hi")) // "Hi"

// Choosing between values
config := lang.Check("", "default")     // "default" (first is empty)
timeout := lang.Check(0, 30)            // 30 (first is zero)
```

### Conditional Operations

Clean conditional logic without if statements:

```go
// Conditional values
message := lang.If(user.IsAdmin, "Admin Panel", "User Panel")

// Conditional function execution
lang.IfF(user.IsLoggedIn, func() {
    log.Println("User logged in")
}, func() {
    log.Println("Anonymous user")
})

// Execute function only if value is non-zero
lang.IfV(user.ID, func() {
    updateUserStats(user.ID)
})
```

### String Operations

Powerful string manipulation and conversion:

```go
// Convert anything to string with optional length limit
text := lang.String(12345)           // "12345"
short := lang.String("Hello World", 5) // "Hello"

// String truncation with ellipsis
truncated := lang.TruncateString("Long text here", 8, "...") // "Long tex..."

// Path manipulation
path := lang.GetWithSep("config", '/')  // "config/"
path = lang.GetWithSep("config/", '/')  // "config/" (unchanged)
```

## Slice Operations

### Transformation and Filtering

```go
// Transform elements to different types
numbers := []int{1, 2, 3, 4, 5}
strings := lang.Convert(numbers, func(n int) string {
    return fmt.Sprintf("#%d", n)
})
// strings: ["#1", "#2", "#3", "#4", "#5"]

// Filter with complex conditions
users := []User{
    {Name: "Alice", Age: 25, Active: true},
    {Name: "Bob", Age: 17, Active: false},
    {Name: "Charlie", Age: 30, Active: true},
}

activeAdults := lang.Filter(users, func(u User) bool {
    return u.Age >= 18 && u.Active
})
// activeAdults: [Alice, Charlie]

// Transform with error handling
inputs := []string{"1", "2", "invalid", "4"}
numbers, err := lang.ConvertWithErr(inputs, func(s string) (int, error) {
    return strconv.Atoi(s)
})
// numbers: nil, err: strconv.Atoi: parsing "invalid": invalid syntax
```

### Search and Analysis

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Find first match
first, found := lang.FindFirst(numbers, func(n int) bool {
    return n > 5
})
// first: 6, found: true

// Check if all/any elements satisfy condition
allPositive := lang.All(numbers, func(n int) bool { return n > 0 })  // true
hasEven := lang.Any(numbers, func(n int) bool { return n%2 == 0 })   // true

// Find indices
index := lang.IndexOf([]string{"a", "b", "c", "b"}, "b")     // 1
lastIndex := lang.LastIndexOf([]string{"a", "b", "c", "b"}, "b") // 3
```

### Data Manipulation

```go
// Remove duplicates while preserving order
unique := lang.Distinct([]int{1, 2, 2, 3, 1, 4}) // [1, 2, 3, 4]

// Set operations
set1 := []int{1, 2, 3, 4}
set2 := []int{3, 4, 5, 6}

intersection := lang.Intersect(set1, set2)  // [3, 4]
union := lang.Union(set1, set2)             // [1, 2, 3, 4, 5, 6]
difference := lang.Difference(set1, set2)   // [1, 2]

// Chunking and partitioning
chunks := lang.Chunk([]int{1, 2, 3, 4, 5, 6, 7}, 3) // [[1,2,3], [4,5,6], [7]]

evens, odds := lang.Partition([]int{1, 2, 3, 4, 5}, func(n int) bool {
    return n%2 == 0
})
// evens: [2, 4], odds: [1, 3, 5]
```

### Advanced Operations

```go
// Reduce (fold) operation
sum := lang.Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int {
    return acc + n
})
// sum: 15

// Take and skip
first3 := lang.Take([]int{1, 2, 3, 4, 5}, 3)  // [1, 2, 3]
after2 := lang.Skip([]int{1, 2, 3, 4, 5}, 2)   // [3, 4, 5]

// Flatten nested slices
nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
flat := lang.Flatten(nested) // [1, 2, 3, 4, 5, 6]

// Reverse
reversed := lang.Reverse([]int{1, 2, 3, 4, 5}) // [5, 4, 3, 2, 1]
```

## Map Operations

### Creation and Transformation

```go
// Create maps from slices
users := []User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// Map by key function
usersByID := lang.SliceToMapByKey(users, func(u User) int {
    return u.ID
})
// usersByID: map[int]User{1: {ID: 1, Name: "Alice", ...}, 2: {ID: 2, Name: "Bob", ...}}

// Transform to different key-value pairs
emailMap := lang.SliceToMap(users, func(u User) (string, string) {
    return u.Name, u.Email
})
// emailMap: map[string]string{"Alice": "alice@example.com", "Bob": "bob@example.com"}

// Create from pairs
pairs := []string{"name", "John", "age", "25", "city", "NYC"}
config := lang.PairsToMap(pairs)
// config: map[string]string{"name": "John", "age": "25", "city": "NYC"}
```

### Map Manipulation

```go
ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 17}

// Filter maps
adults := lang.FilterMap(ages, func(name string, age int) bool {
    return age >= 18
})
// adults: map[string]int{"Alice": 25, "Bob": 30}

// Transform values
descriptions := lang.ConvertMap(ages, func(age int) string {
    return fmt.Sprintf("%d years old", age)
})
// descriptions: map[string]string{"Alice": "25 years old", "Bob": "30 years old", "Charlie": "17 years old"}

// Extract keys and values
names := lang.Keys(ages)                    // ["Alice", "Bob", "Charlie"]
adultNames := lang.KeysIf(ages, func(name string, age int) bool {
    return age >= 18
})                                         // ["Alice", "Bob"]
```

### Grouping and Aggregation

```go
// Group by category
products := []Product{
    {Name: "iPhone", Category: "Electronics", Price: 999},
    {Name: "iPad", Category: "Electronics", Price: 799},
    {Name: "Book", Category: "Literature", Price: 25},
    {Name: "Pen", Category: "Office", Price: 5},
}

byCategory := lang.GroupBy(products, func(p Product) string {
    return p.Category
})
// byCategory: map[string][]Product{
//   "Electronics": [{iPhone...}, {iPad...}],
//   "Literature": [{Book...}],
//   "Office": [{Pen...}],
// }

// Merge multiple maps
priceMap1 := map[string]int{"apple": 2, "banana": 1}
priceMap2 := map[string]int{"banana": 3, "orange": 4}
merged := lang.MergeMap(priceMap1, priceMap2)
// merged: map[string]int{"apple": 2, "banana": 3, "orange": 4}
```

## Error Handling

### Panic Recovery

```go
// Automatic goroutine recovery with restart
logger := &MyLogger{}
lang.Go(logger, func() {
    // This goroutine will restart if it panics
    for {
        riskyOperation()
        time.Sleep(time.Second)
    }
})

// Function-level panic recovery
func riskyFunction() (err error) {
    defer func() {
        if lang.RecoverWithErr(&err) {
            // Panic was converted to error
        }
    }()
    
    // Code that might panic
    return nil
}

// Default value on panic
result := lang.DefaultIfPanic("safe default", func() string {
    return mightPanicFunction()
})
```

### Error Utilities

```go
// Wrap errors with context
err := someOperation()
if err != nil {
    return lang.Wrap(err, "failed to perform operation")
}

// Combine multiple errors
err1 := operation1()
err2 := operation2()
if combinedErr := lang.JoinErrors(err1, err2); combinedErr != nil {
    return combinedErr
}
```

## Type Safety

### Type Conversion and Checking

```go
// Safe type conversion
var value any = "hello"
str := lang.Type[string](value)  // "hello"
num := lang.Type[int](value)     // 0 (zero value, conversion failed)

// Retry mechanism
result, err := lang.Retry(3, func() (string, error) {
    return callExternalAPI()
})

// Timeout handling
result, err := lang.RunWithTimeout(5*time.Second, func() (string, error) {
    return longRunningOperation()
})
```

## Real-World Examples

### Data Processing Pipeline

```go
// Process user data with validation and transformation
func processUserData(rawData []map[string]interface{}) []UserProfile {
    // Convert raw data to structured format
    users := lang.Convert(rawData, func(raw map[string]interface{}) User {
        return User{
            ID:    lang.Type[int](raw["id"]),
            Name:  lang.Type[string](raw["name"]),
            Email: lang.Type[string](raw["email"]),
            Age:   lang.Type[int](raw["age"]),
        }
    })
    
    // Filter valid users
    validUsers := lang.Filter(users, func(u User) bool {
        return u.ID > 0 && u.Name != "" && strings.Contains(u.Email, "@")
    })
    
    // Group by age range
    byAgeGroup := lang.GroupBy(validUsers, func(u User) string {
        switch {
        case u.Age < 18:
            return "minor"
        case u.Age < 65:
            return "adult"
        default:
            return "senior"
        }
    })
    
    // Create user profiles
    profiles := lang.ConvertFromMap(byAgeGroup, func(group string, users []User) UserProfile {
        return UserProfile{
            AgeGroup: group,
            Count:    len(users),
            Users:    users,
        }
    })
    
    return profiles
}
```

### Configuration Management

```go
// Safe configuration handling
type Config struct {
    Database DatabaseConfig
    Server   ServerConfig
    Features map[string]bool
}

func loadConfig() Config {
    // Load with defaults
    return Config{
        Database: DatabaseConfig{
            Host:     lang.Check(os.Getenv("DB_HOST"), "localhost"),
            Port:     lang.Check(parsePort(os.Getenv("DB_PORT")), 5432),
            Timeout:  lang.Check(parseDuration(os.Getenv("DB_TIMEOUT")), 30*time.Second),
        },
        Server: ServerConfig{
            Port:    lang.Check(parsePort(os.Getenv("SERVER_PORT")), 8080),
            Workers: lang.Check(parseWorkers(os.Getenv("WORKERS")), 4),
        },
        Features: lang.PairsToMap(lang.NotEmpty(strings.Split(os.Getenv("FEATURES"), ","))),
    }
}
```

### API Response Processing

```go
// Process API responses with error handling
func processAPIResponse(responses []APIResponse) ([]ProcessedData, error) {
    // Filter successful responses
    successful := lang.Filter(responses, func(r APIResponse) bool {
        return r.Status == "success"
    })
    
    // Extract and process data
    data, err := lang.ConvertWithErr(successful, func(r APIResponse) (ProcessedData, error) {
        return ProcessedData{
            ID:        r.ID,
            Value:     r.Data.Value,
            Timestamp: time.Now(),
        }, nil
    })
    
    if err != nil {
        return nil, lang.Wrap(err, "failed to process API responses")
    }
    
    // Remove duplicates and sort
    unique := lang.Distinct(data)
    
    return unique, nil
}
```

## Performance Notes

- **Memory Allocation**: Most functions pre-allocate slices and maps with appropriate capacity to minimize allocations
- **Goroutine Safety**: All functions are goroutine-safe for read operations. Map operations are not goroutine-safe for concurrent writes
- **Large Data**: For very large datasets (>1M elements), consider processing in chunks using `Chunk()` function
- **Error Handling**: Functions with error handling (`ConvertWithErr`, `RecoverWithErr`) have minimal overhead when no errors occur

## Best Practices

### Function Composition

```go
// Chain operations for cleaner code
result := lang.Filter(
    lang.Convert(rawData, parseUser),
    func(u User) bool { return u.IsValid() },
)

// Use with method chaining
type Pipeline[T any] struct {
    data []T
}

func (p Pipeline[T]) Filter(fn func(T) bool) Pipeline[T] {
    return Pipeline[T]{lang.Filter(p.data, fn)}
}

func (p Pipeline[T]) Map(fn func(T) T) Pipeline[T] {
    return Pipeline[T]{lang.Map(p.data, fn)}
}
```

### Error Handling Strategy

```go
// Prefer error returns over panics
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Use panic recovery only for truly exceptional cases
func mustParseConfig(filename string) Config {
    defer func() {
        if lang.Recover(logger) {
            log.Fatal("Failed to parse critical configuration")
        }
    }()
    
    return parseConfig(filename)
}
```

### Memory Management

```go
// For large slices, use copy operations to avoid memory leaks
func processLargeSlice(input []LargeStruct) []Result {
    // Process only what you need
    relevant := lang.Filter(input, isRelevant)
    
    // Copy to avoid holding reference to original large slice
    return lang.Convert(lang.Copy(relevant), processStruct)
}
```

## Comparison with Native Go

**Without lang** (traditional Go):
```go
// Filtering
var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}

// Mapping
var doubled []int
for _, n := range numbers {
    doubled = append(doubled, n*2)
}

// Grouping
groups := make(map[string][]Person)
for _, person := range people {
    key := person.Department
    groups[key] = append(groups[key], person)
}
```

**With lang** (functional approach):
```go
// Filtering
evens := lang.Filter(numbers, func(n int) bool { return n%2 == 0 })

// Mapping
doubled := lang.Map(numbers, func(n int) int { return n * 2 })

// Grouping
groups := lang.GroupBy(people, func(p Person) string { return p.Department })
```

## Related Libraries

- **[lo](https://github.com/samber/lo)**: Another excellent functional programming library for Go
- **[go-funk](https://github.com/thoas/go-funk)**: Functional utilities using reflection
- **[pie](https://github.com/elliotchance/pie)**: Type-safe slice operations

---

**Add functional programming power to your Go projects: [Documentation][doc]**

[version-img]: https://img.shields.io/badge/Go-%3E%3D%201.19-%23007d9c
[doc-img]: https://pkg.go.dev/badge/github.com/maxbolgarin/lang
[doc]: https://pkg.go.dev/github.com/maxbolgarin/lang
[ci-img]: https://github.com/maxbolgarin/lang/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/maxbolgarin/lang/actions
[report-img]: https://goreportcard.com/badge/github.com/maxbolgarin/lang
[report]: https://goreportcard.com/report/github.com/maxbolgarin/lang
[coverage-img]: https://codecov.io/gh/maxbolgarin/lang/branch/main/graph/badge.svg
[coverage]: https://codecov.io/gh/maxbolgarin/lang
