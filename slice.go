package lang

// SliceToMap creates a map by transforming each element of a slice into a key-value pair.
// The transform function should return a key and value for each element.
//
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	userMap := SliceToMap(users, func(u User) (int, string) {
//	    return u.ID, u.Name
//	}) // userMap == map[int]string{1: "Alice", 2: "Bob"}
func SliceToMap[T any, K comparable, V any](input []T, transform func(T) (K, V)) map[K]V {
	if input == nil {
		return make(map[K]V)
	}
	out := make(map[K]V, len(input))
	for _, e := range input {
		k, v := transform(e)
		out[k] = v
	}
	return out
}

// SliceToMapByKey creates a map by using a key function to generate keys from slice elements.
// The elements themselves become the values in the resulting map.
//
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	userMap := SliceToMapByKey(users, func(u User) int {
//	    return u.ID
//	}) // userMap == map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
func SliceToMapByKey[T any, K comparable](input []T, key func(T) K) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	return SliceToMap(input, func(t T) (K, T) { return key(t), t })
}

// Mapping is an alias for SliceToMapByKey that creates a map by using a key function.
//
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	userMap := Mapping(users, func(u User) int {
//	    return u.ID
//	}) // userMap == map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
func Mapping[T any, K comparable](input []T, key func(T) K) map[K]T {
	return SliceToMapByKey(input, key)
}

// PairsToMap transforms a slice with pairs of elements into a map.
// The first element of each pair becomes the key, and the second becomes the value.
// If the slice has an odd number of elements, the last element is ignored.
//
//	pairs := []string{"key1", "value1", "key2", "value2", "key3"}
//	mapping := PairsToMap(pairs) // mapping == map[string]string{"key1": "value1", "key2": "value2"}
func PairsToMap[T comparable](input []T) map[T]T {
	if input == nil {
		return make(map[T]T)
	}
	out := make(map[T]T, len(input)/2)
	for i := 0; i < len(input)-1; i += 2 {
		out[input[i]] = input[i+1]
	}
	return out
}

// Filter returns a new slice containing only the elements that satisfy the filter function.
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := Filter(numbers, func(n int) bool { return n%2 == 0 }) // evens == []int{2, 4, 6}
func Filter[T any](input []T, filter func(T) bool) []T {
	if input == nil {
		return nil
	}
	out := make([]T, 0, len(input))
	for _, e := range input {
		if filter(e) {
			out = append(out, e)
		}
	}
	return out
}

// Map transforms each element of a slice using the provided function and returns a new slice.
//
//	numbers := []int{1, 2, 3}
//	doubled := Map(numbers, func(n int) int { return n * 2 }) // doubled == []int{2, 4, 6}
func Map[T any](input []T, transform func(T) T) []T {
	if input == nil {
		return nil
	}
	out := make([]T, 0, len(input))
	for _, e := range input {
		out = append(out, transform(e))
	}
	return out
}

// Reduce applies a function to each element in a slice to reduce it to a single value.
//
//	nums := []int{1, 2, 3, 4}
//	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n }) // sum == 10
func Reduce[T, K any](s []T, initial K, f func(K, T) K) K {
	if s == nil {
		return initial
	}
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

// Convert transforms each element of a slice to a different type using the provided function.
//
//	numbers := []int{1, 2, 3}
//	strings := Convert(numbers, func(n int) string { return strconv.Itoa(n) }) // strings == []string{"1", "2", "3"}
func Convert[T, K any](input []T, transform func(T) K) []K {
	if input == nil {
		return nil
	}
	out := make([]K, 0, len(input))
	for _, e := range input {
		out = append(out, transform(e))
	}
	return out
}

// ConvertWithErr transforms each element of a slice to a different type using the provided function.
// Returns an error if any transformation fails.
//
//	strings := []string{"1", "2", "invalid"}
//	numbers, err := ConvertWithErr(strings, func(s string) (int, error) {
//	    return strconv.Atoi(s)
//	}) // numbers == nil, err != nil
func ConvertWithErr[T, K any](input []T, transform func(T) (K, error)) ([]K, error) {
	if input == nil {
		return nil, nil
	}
	out := make([]K, 0, len(input))
	for _, e := range input {
		res, err := transform(e)
		if err != nil {
			return nil, err
		}
		out = append(out, res)
	}
	return out, nil
}

// ConvertMap transforms each value in a map using the provided function while preserving keys.
//
//	ages := map[string]int{"Alice": 25, "Bob": 30}
//	descriptions := ConvertMap(ages, func(age int) string {
//	    return fmt.Sprintf("%d years old", age)
//	}) // descriptions == map[string]string{"Alice": "25 years old", "Bob": "30 years old"}
func ConvertMap[K comparable, T1, T2 any](input map[K]T1, transform func(T1) T2) map[K]T2 {
	if input == nil {
		return make(map[K]T2)
	}
	out := make(map[K]T2, len(input))
	for k, v := range input {
		out[k] = transform(v)
	}
	return out
}

// ConvertMapWithErr transforms each value in a map using the provided function while preserving keys.
// Returns an error if any transformation fails.
//
//	stringNums := map[string]string{"a": "1", "b": "invalid"}
//	numbers, err := ConvertMapWithErr(stringNums, func(s string) (int, error) {
//	    return strconv.Atoi(s)
//	}) // numbers == nil, err != nil
func ConvertMapWithErr[K comparable, T1, T2 any](input map[K]T1, transform func(T1) (T2, error)) (map[K]T2, error) {
	if input == nil {
		return make(map[K]T2), nil
	}
	out := make(map[K]T2, len(input))
	for k, v := range input {
		res, err := transform(v)
		if err != nil {
			return nil, err
		}
		out[k] = res
	}
	return out, nil
}

// ConvertFromMap transforms each key-value pair in a map into a slice element.
//
//	ages := map[string]int{"Alice": 25, "Bob": 30}
//	descriptions := ConvertFromMap(ages, func(name string, age int) string {
//	    return fmt.Sprintf("%s is %d years old", name, age)
//	}) // descriptions == []string{"Alice is 25 years old", "Bob is 30 years old"}
func ConvertFromMap[K comparable, T1, T2 any](input map[K]T1, transform func(K, T1) T2) []T2 {
	if input == nil {
		return nil
	}
	out := make([]T2, 0, len(input))
	for k, v := range input {
		out = append(out, transform(k, v))
	}
	return out
}

// ConvertFromMapWithErr transforms each key-value pair in a map into a slice element.
// Returns an error if any transformation fails.
//
//	stringNums := map[string]string{"a": "1", "b": "invalid"}
//	numbers, err := ConvertFromMapWithErr(stringNums, func(key string, val string) (int, error) {
//	    return strconv.Atoi(val)
//	}) // numbers == nil, err != nil
func ConvertFromMapWithErr[K comparable, T1, T2 any](input map[K]T1, transform func(K, T1) (T2, error)) ([]T2, error) {
	if input == nil {
		return nil, nil
	}
	out := make([]T2, 0, len(input))
	for k, v := range input {
		res, err := transform(k, v)
		if err != nil {
			return nil, err
		}
		out = append(out, res)
	}
	return out, nil
}

// ConvertToMap transforms each element of a slice into a key-value pair for a map.
//
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	userMap := ConvertToMap(users, func(u User) (int, string) {
//	    return u.ID, u.Name
//	}) // userMap == map[int]string{1: "Alice", 2: "Bob"}
func ConvertToMap[T1 any, K comparable, T2 any](input []T1, transform func(T1) (K, T2)) map[K]T2 {
	if input == nil {
		return make(map[K]T2)
	}
	out := make(map[K]T2, len(input))
	for _, e := range input {
		k, v := transform(e)
		out[k] = v
	}
	return out
}

// ConvertToMapWithErr transforms each element of a slice into a key-value pair for a map.
// Returns an error if any transformation fails.
//
//	strings := []string{"1", "2", "invalid"}
//	numberMap, err := ConvertToMapWithErr(strings, func(s string) (string, int, error) {
//	    num, err := strconv.Atoi(s)
//	    return s, num, err
//	}) // numberMap == nil, err != nil
func ConvertToMapWithErr[T1 any, K comparable, T2 any](input []T1, transform func(T1) (K, T2, error)) (map[K]T2, error) {
	if input == nil {
		return make(map[K]T2), nil
	}
	out := make(map[K]T2, len(input))
	for _, e := range input {
		k, v, err := transform(e)
		if err != nil {
			return nil, err
		}
		out[k] = v
	}
	return out, nil
}

// FilterMap returns a new map containing only the key-value pairs that satisfy the filter function.
//
//	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 17}
//	adults := FilterMap(ages, func(name string, age int) bool {
//	    return age >= 18
//	}) // adults == map[string]int{"Alice": 25, "Bob": 30}
func FilterMap[K comparable, T any](input map[K]T, filter func(K, T) bool) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	out := make(map[K]T, len(input))
	for k, v := range input {
		if filter(k, v) {
			out[k] = v
		}
	}
	return out
}

// Copy creates a shallow copy of a slice.
//
//	original := []int{1, 2, 3}
//	copied := Copy(original) // copied == []int{1, 2, 3}
//	original[0] = 99         // copied[0] is still 1
func Copy[T any](input []T) []T {
	if input == nil {
		return nil
	}
	out := make([]T, len(input))
	copy(out, input)
	return out
}

// CopyMap creates a shallow copy of a map.
//
//	original := map[string]int{"a": 1, "b": 2}
//	copied := CopyMap(original) // copied == map[string]int{"a": 1, "b": 2}
//	original["a"] = 99          // copied["a"] is still 1
func CopyMap[K comparable, T any](input map[K]T) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	out := make(map[K]T, len(input))
	for k, v := range input {
		out[k] = v
	}
	return out
}

// WithoutEmpty returns a new slice with all zero values removed.
//
//	mixed := []string{"hello", "", "world", ""}
//	nonEmpty := WithoutEmpty(mixed) // nonEmpty == []string{"hello", "world"}
func WithoutEmpty[T comparable](input []T) []T {
	if input == nil {
		return nil
	}
	var empty T
	out := make([]T, 0, len(input))
	for _, v := range input {
		if v != empty {
			out = append(out, v)
		}
	}
	return out
}

// Keys returns a slice containing all keys from a map.
//
//	mapping := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := Keys(mapping) // keys == []string{"a", "b", "c"} (order may vary)
func Keys[K comparable, T any](input map[K]T) []K {
	if input == nil {
		return nil
	}
	out := make([]K, 0, len(input))
	for k := range input {
		out = append(out, k)
	}
	return out
}

// KeysIf returns a slice containing keys from a map that satisfy the filter function.
//
//	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 17}
//	adultNames := KeysIf(ages, func(name string, age int) bool {
//	    return age >= 18
//	}) // adultNames == []string{"Alice", "Bob"} (order may vary)
func KeysIf[K comparable, T any](input map[K]T, filter func(K, T) bool) []K {
	if input == nil {
		return nil
	}
	out := make([]K, 0, len(input))
	for k, v := range input {
		if !filter(k, v) {
			continue
		}
		out = append(out, k)
	}
	return out
}

// Values returns a slice containing all values from a map.
//
//	mapping := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := Values(mapping) // values == []int{1, 2, 3} (order may vary)
func Values[K comparable, T any](input map[K]T) []T {
	if input == nil {
		return nil
	}
	out := make([]T, 0, len(input))
	for _, v := range input {
		out = append(out, v)
	}
	return out
}

// ValuesIf returns a slice containing values from a map that satisfy the filter function.
//
//	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 17}
//	adultAges := ValuesIf(ages, func(name string, age int) bool {
//	    return age >= 18
//	}) // adultAges == []int{25, 30} (order may vary)
func ValuesIf[K comparable, T any](input map[K]T, filter func(K, T) bool) []T {
	if input == nil {
		return nil
	}
	out := make([]T, 0, len(input))
	for k, v := range input {
		if !filter(k, v) {
			continue
		}
		out = append(out, v)
	}
	return out
}

// WithoutEmptyKeys returns a new map with all entries that have zero-value keys removed.
//
//	mapping := map[string]int{"": 1, "a": 2, "b": 3}
//	filtered := WithoutEmptyKeys(mapping) // filtered == map[string]int{"a": 2, "b": 3}
func WithoutEmptyKeys[K comparable, T any](input map[K]T) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	var empty K
	out := make(map[K]T, len(input))
	for k, v := range input {
		if k != empty {
			out[k] = v
		}
	}
	return out
}

// WithoutEmptyValues returns a new map with all entries that have zero-value values removed.
//
//	mapping := map[string]int{"a": 0, "b": 2, "c": 3}
//	filtered := WithoutEmptyValues(mapping) // filtered == map[string]int{"b": 2, "c": 3}
func WithoutEmptyValues[K, T comparable](input map[K]T) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	var empty T
	out := make(map[K]T, len(input))
	for k, v := range input {
		if v != empty {
			out[k] = v
		}
	}
	return out
}

// NotEmpty is an alias for WithoutEmpty that returns a new slice with all zero values removed.
//
//	mixed := []string{"hello", "", "world", ""}
//	nonEmpty := NotEmpty(mixed) // nonEmpty == []string{"hello", "world"}
func NotEmpty[T comparable](input []T) []T {
	return WithoutEmpty(input)
}

// NotEmptyMap returns a new map with all entries that have zero-value keys or values removed.
//
//	mapping := map[string]int{"": 1, "a": 0, "b": 2}
//	filtered := NotEmptyMap(mapping) // filtered == map[string]int{"b": 2}
func NotEmptyMap[K, T comparable](input map[K]T) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	var empty T
	var emptyK K
	out := make(map[K]T, len(input))
	for k, v := range input {
		if v != empty && k != emptyK {
			out[k] = v
		}
	}
	return out
}

// SplitByChunkSize splits a slice into chunks of the specified size.
// If chunkSize is less than 1, it will be treated as 1.
//
//	items := []int{1, 2, 3, 4, 5, 6, 7}
//	chunks := SplitByChunkSize(items, 3) // chunks == [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
func SplitByChunkSize[T any](items []T, chunkSize int) [][]T {
	if items == nil {
		return nil
	}

	if chunkSize < 1 {
		chunkSize = 1
	}

	var chunks [][]T
	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}

	return chunks
}

// FindFirst returns the first element in a slice that satisfies a predicate.
// Returns zero value and false if no element is found.
//
//	nums := []int{1, 2, 3, 4, 5}
//	first, found := FindFirst(nums, func(n int) bool { return n > 3 }) // first == 4, found == true
func FindFirst[T any](s []T, predicate func(T) bool) (T, bool) {
	if s == nil {
		var zero T
		return zero, false
	}

	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Contains checks if a slice contains a specific element.
//
//	found := Contains([]int{1, 2, 3}, 2) // found == true
func Contains[T comparable](s []T, element T) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == element {
			return true
		}
	}
	return false
}

// ContainsFunc checks if a slice contains an element that satisfies the predicate.
//
//	found := ContainsFunc([]int{1, 2, 3}, func(n int) bool { return n > 2 }) // found == true
func ContainsFunc[T any](s []T, predicate func(T) bool) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if predicate(v) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of an element in a slice, or -1 if not found.
//
//	idx := IndexOf([]string{"a", "b", "c"}, "b") // idx == 1
func IndexOf[T comparable](s []T, element T) int {
	if s == nil {
		return -1
	}
	for i, v := range s {
		if v == element {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of an element in a slice, or -1 if not found.
//
//	idx := LastIndexOf([]string{"a", "b", "c", "b"}, "b") // idx == 3
func LastIndexOf[T comparable](s []T, element T) int {
	if s == nil {
		return -1
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == element {
			return i
		}
	}
	return -1
}

// Distinct returns a new slice with duplicate elements removed, preserving order.
//
//	unique := Distinct([]int{1, 2, 2, 3, 1, 4}) // unique == []int{1, 2, 3, 4}
func Distinct[T comparable](s []T) []T {
	if s == nil {
		return nil
	}
	result := make([]T, 0, len(s))
	seen := make(map[T]struct{}, len(s))
	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Intersect returns elements that exist in both slices.
//
//	common := Intersect([]int{1, 2, 3}, []int{2, 3, 4}) // common == []int{2, 3}
func Intersect[T comparable](a, b []T) []T {
	if a == nil || b == nil {
		return nil
	}

	if len(a) == 0 || len(b) == 0 {
		return []T{}
	}

	// Use the smaller slice to build the lookup map for efficiency
	var lookup map[T]struct{}
	var iterate []T

	if len(a) <= len(b) {
		lookup = make(map[T]struct{}, len(a))
		for _, v := range a {
			lookup[v] = struct{}{}
		}
		iterate = b
	} else {
		lookup = make(map[T]struct{}, len(b))
		for _, v := range b {
			lookup[v] = struct{}{}
		}
		iterate = a
	}

	// Build result with elements that exist in both
	result := make([]T, 0)
	seen := make(map[T]struct{})

	for _, v := range iterate {
		if _, exists := lookup[v]; exists {
			if _, alreadySeen := seen[v]; !alreadySeen {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}

	return result
}

// Union returns a slice with unique elements from all input slices.
//
//	all := Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{4, 5, 6}) // all == []int{1, 2, 3, 4, 5, 6}
func Union[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	// Check if all slices are nil
	allNil := true
	for _, s := range slices {
		if s != nil {
			allNil = false
			break
		}
	}
	if allNil {
		return nil
	}

	// Calculate total capacity for worst case
	totalCap := 0
	for _, s := range slices {
		if s != nil {
			totalCap += len(s)
		}
	}

	// Collect all unique elements
	result := make([]T, 0, totalCap)
	seen := make(map[T]struct{}, totalCap)

	for _, slice := range slices {
		if slice == nil {
			continue
		}

		for _, v := range slice {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}

	return result
}

// Difference returns elements that exist in the first slice but not in the second.
//
//	diff := Difference([]int{1, 2, 3, 4}, []int{2, 4}) // diff == []int{1, 3}
func Difference[T comparable](a, b []T) []T {
	if a == nil {
		return nil
	}
	if len(b) == 0 {
		return Copy(a)
	}

	// Build lookup of elements to exclude
	exclude := make(map[T]struct{}, len(b))
	for _, v := range b {
		exclude[v] = struct{}{}
	}

	// Add elements from first slice that aren't in the second
	result := make([]T, 0, len(a))
	for _, v := range a {
		if _, exists := exclude[v]; !exists {
			result = append(result, v)
		}
	}

	return result
}

// Reverse returns a new slice with elements in reverse order.
//
//	reversed := Reverse([]int{1, 2, 3}) // reversed == []int{3, 2, 1}
func Reverse[T any](s []T) []T {
	if s == nil {
		return nil
	}

	result := make([]T, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		result[i] = s[j]
	}

	return result
}

// Flatten transforms a slice of slices into a single slice with all elements.
//
//	flat := Flatten([][]int{{1, 2}, {3, 4}}) // flat == []int{1, 2, 3, 4}
func Flatten[T any](s [][]T) []T {
	if s == nil {
		return nil
	}

	// Calculate total capacity needed
	totalLen := 0
	for _, v := range s {
		totalLen += len(v)
	}

	result := make([]T, 0, totalLen)
	for _, v := range s {
		result = append(result, v...)
	}

	return result
}

// Chunk splits a slice into chunks of the specified size (alias for SplitByChunkSize).
//
//	chunks := Chunk([]int{1, 2, 3, 4, 5}, 2) // chunks == [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[T any](s []T, size int) [][]T {
	return SplitByChunkSize(s, size)
}

// GroupBy groups slice elements by a key generated from each element.
//
//	people := []struct{Name string; Age int}{{"Alice", 25}, {"Bob", 30}, {"Charlie", 25}}
//	byAge := GroupBy(people, func(p struct{Name string; Age int}) int { return p.Age })
//	// byAge == map[int][]struct{Name string; Age int}{
//	//   25: {{"Alice", 25}, {"Charlie", 25}},
//	//   30: {{"Bob", 30}},
//	// }
func GroupBy[T any, K comparable](s []T, keyFn func(T) K) map[K][]T {
	if s == nil {
		return make(map[K][]T)
	}

	result := make(map[K][]T)
	for _, v := range s {
		key := keyFn(v)
		result[key] = append(result[key], v)
	}

	return result
}

// ForEach executes a function for each element in a slice.
//
//	sum := 0
//	ForEach([]int{1, 2, 3}, func(n int) { sum += n })
//	// sum == 6
func ForEach[T any](s []T, f func(T)) {
	if s == nil {
		return
	}

	for _, v := range s {
		f(v)
	}
}

// All returns true if all elements in the slice satisfy the predicate.
//
//	allPositive := All([]int{1, 2, 3}, func(n int) bool { return n > 0 }) // allPositive == true
func All[T any](s []T, predicate func(T) bool) bool {
	if len(s) == 0 {
		return true // Vacuously true
	}

	for _, v := range s {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Any returns true if at least one element in the slice satisfies the predicate.
//
//	hasNegative := Any([]int{1, -2, 3}, func(n int) bool { return n < 0 }) // hasNegative == true
func Any[T any](s []T, predicate func(T) bool) bool {
	return ContainsFunc(s, predicate)
}

// Take returns a slice with the first n elements. If n is greater than the length of the slice,
// the entire slice is returned.
//
//	first3 := Take([]int{1, 2, 3, 4, 5}, 3) // first3 == []int{1, 2, 3}
func Take[T any](s []T, n int) []T {
	if s == nil {
		return nil
	}

	if n <= 0 {
		return []T{}
	}

	if n >= len(s) {
		return Copy(s)
	}

	result := make([]T, n)
	copy(result, s[:n])

	return result
}

// Skip returns a slice without the first n elements. If n is greater than the length of the slice,
// an empty slice is returned.
//
//	rest := Skip([]int{1, 2, 3, 4, 5}, 2) // rest == []int{3, 4, 5}
func Skip[T any](s []T, n int) []T {
	if s == nil {
		return nil
	}

	if n <= 0 {
		return Copy(s)
	}

	if n >= len(s) {
		return []T{}
	}

	result := make([]T, len(s)-n)
	copy(result, s[n:])

	return result
}

// Compact removes nil values from a slice of pointers or interfaces.
//
//	a, b, c := 1, 2, 3
//	ptrs := []*int{&a, nil, &b, nil, &c}
//	nonNil := Compact(ptrs) // nonNil == []*int{&a, &b, &c}
func Compact[T any](s []*T) []*T {
	if s == nil {
		return nil
	}

	result := make([]*T, 0, len(s))
	for _, v := range s {
		if v != nil {
			result = append(result, v)
		}
	}

	return result
}

// MergeMap merges multiple maps into a single map. In case of key conflicts,
// values from later maps overwrite earlier ones.
//
//	merged := MergeMap(
//	    map[string]int{"a": 1, "b": 2},
//	    map[string]int{"b": 3, "c": 4},
//	) // merged == map[string]int{"a": 1, "b": 3, "c": 4}
func MergeMap[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return make(map[K]V)
	}

	// Count total capacity needed
	totalSize := 0
	for _, m := range maps {
		if m != nil {
			totalSize += len(m)
		}
	}

	result := make(map[K]V, totalSize)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// ZipToMap creates a map from two slices, using the first slice for keys and the second for values.
// If the slices have different lengths, the extra elements from the longer slice are ignored.
//
//	keys := []string{"a", "b", "c"}
//	values := []int{1, 2, 3}
//	mapping := ZipToMap(keys, values) // mapping == map[string]int{"a": 1, "b": 2, "c": 3}
func ZipToMap[K comparable, V any](keys []K, values []V) map[K]V {
	if keys == nil || values == nil {
		return make(map[K]V)
	}

	minLen := len(keys)
	if len(values) < minLen {
		minLen = len(values)
	}

	result := make(map[K]V, minLen)
	for i := 0; i < minLen; i++ {
		result[keys[i]] = values[i]
	}

	return result
}

// Partition splits a slice into two slices based on a predicate function.
// The first slice contains elements that satisfy the predicate, the second contains those that don't.
//
//	evens, odds := Partition([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
//	// evens == []int{2, 4}, odds == []int{1, 3, 5}
func Partition[T any](s []T, predicate func(T) bool) ([]T, []T) {
	if s == nil {
		return nil, nil
	}

	matching := make([]T, 0, len(s)/2)
	nonMatching := make([]T, 0, len(s)/2)

	for _, v := range s {
		if predicate(v) {
			matching = append(matching, v)
		} else {
			nonMatching = append(nonMatching, v)
		}
	}

	return matching, nonMatching
}

// TruncateSlice truncates a slice to a maximum length.
// It is not change capacity of the slice, so items will be still in the underlying array.
//
//	a := []int{1, 2, 3}
//	b := TruncateSlice(a, 2) // b == []int{1, 2}
func TruncateSlice[T any](s []T, maxLen int) []T {
	if s == nil {
		return nil
	}
	if maxLen <= 0 {
		return s[:0]
	}
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// TruncateSliceWithCopy truncates a slice to a maximum length and returns a new slice.
// Old slice will be garbage collected if there are no references to it.
//
//	a := []int{1, 2, 3}
//	b := TruncateSliceWithCopy(a, 2) // b == []int{1, 2}
func TruncateSliceWithCopy[T any](s []T, maxLen int) []T {
	if s == nil {
		return nil
	}
	if maxLen <= 0 {
		return s[:0]
	}
	if len(s) <= maxLen {
		copied := make([]T, len(s))
		copy(copied, s)
		return copied
	}
	copied := make([]T, maxLen)
	copy(copied, s[:maxLen])
	return copied
}

// Slice returns a slice of the given type.
// If the input is a slice, it is truncated to the given length.
// If the input is a single value, it is returned as a slice of length 1.
// If the input is not a slice or a single value, it is returned as nil.
//
//	a := []int{1, 2, 3}
//	b := Slice(a, 2) // b == []int{1, 2}
func Slice[T any](s any, maxLenRaw ...int) []T {
	if s == nil {
		return nil
	}
	var maxLen int
	if len(maxLenRaw) > 0 {
		maxLen = maxLenRaw[0]
		if maxLen <= 0 {
			return []T{}
		}
	}
	slice, ok := s.([]T)
	if ok {
		return TruncateSlice(slice, Check(maxLen, len(slice)))
	}
	val, ok := s.(T)
	if ok {
		return []T{val}
	}
	return nil
}
