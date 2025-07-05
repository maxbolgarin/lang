package lang

// SliceToMap returns a new map created calling a transform function on every element of slice,
// function returns a key and an according value. Return empty key to pass iteration.
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

// SliceToMapByKey returns a new map created calling a transform function on every element of slice,
// function returns a key and current element of slice becomes a value.
func SliceToMapByKey[T any, K comparable](input []T, key func(T) K) map[K]T {
	if input == nil {
		return make(map[K]T)
	}
	return SliceToMap(input, func(t T) (K, T) { return key(t), t })
}

// Mapping returns a new map created calling a transform function on every element of slice,
// function returns a key and current element of slice becomes a value.
func Mapping[T any, K comparable](input []T, key func(T) K) map[K]T {
	return SliceToMapByKey(input, key)
}

// PairsToMap transforms a slice with pairs of elements into a map.
// The first element of each pair is a key and the second is a value.
// If the input slice has an odd number of elements, the last element is ignored.
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

// Filter returns a new slice with elements filtered by the given filter function.
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

// Map returns a new slice with elements transformed by the given function with the same type.
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

// Convert returns a new slice with elements transformed by the given function with another type.
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

// ConvertWithErr returns a new slice with elements transformed by the given function with another type.
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

// ConvertMap returns a new map with elements transformed by the given function with another type.
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

// ConvertMapWithErr returns a new map with elements transformed by the given function with another type.
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

// ConvertFromMap returns a new slice with elements transformed by the given function with another type.
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

// ConvertFromMapWithErr returns a new slice with elements transformed by the given function with another type.
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

// ConvertToMap returns a new map with elements transformed by the given function with another type.
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

// ConvertToMapWithErr returns a new map with elements transformed by the given function with another type.
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

// FilterMap returns a new map with elements filtered by the given filter function.
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

// Copy returns a copy of a provided slice.
func Copy[T any](input []T) []T {
	if input == nil {
		return nil
	}
	out := make([]T, len(input))
	copy(out, input)
	return out
}

// CopyMap returns a copy of a provided map.
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

// WithoutEmpty returns a new slice without empty elements.
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

// Keys returns a new slice with keys of a provided map.
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

// KeysIf returns a new slice with keys of a provided map filtered by the given filter function.
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

// Values returns a new slice with values of a provided map.
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

// ValuesIf returns a new slice with values of a provided map filtered by the given filter function.
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

// WithoutEmptyKeys returns a new map without empty keys.
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

// WithoutEmptyValues returns a new map without empty values.
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

// NotEmpty returns a new slice without empty elements.
func NotEmpty[T comparable](input []T) []T {
	return WithoutEmpty(input)
}

// NotEmptyMap returns a new map without empty keys and values.
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
