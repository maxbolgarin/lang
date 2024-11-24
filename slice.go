package lang

// SliceToMap returns a new map created calling a transform function on every element of slice,
// function returns a key and an according value. Return empty key to pass iteration.
func SliceToMap[T any, K comparable, V any](input []T, transform func(T) (K, V)) map[K]V {
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
	return SliceToMap(input, func(t T) (K, T) { return key(t), t })
}

// PairsToMap transforms a slice with pairs of elements into a map.
// The first element of each pair is a key and the second is a value.
func PairsToMap[T comparable](input []T) map[T]T {
	out := make(map[T]T, len(input)/2)
	for i := 0; i < len(input)-1; i += 2 {
		out[input[i]] = input[i+1]
	}
	return out
}

// Filter returns a new slice with elements filtered by the given filter function.
func Filter[T any](input []T, filter func(T) bool) []T {
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
	out := make([]T, 0, len(input))
	for _, e := range input {
		out = append(out, transform(e))
	}
	return out
}

// Convert returns a new slice with elements transformed by the given function with another type.
func Convert[T, K any](input []T, transform func(T) K) []K {
	out := make([]K, 0, len(input))
	for _, e := range input {
		out = append(out, transform(e))
	}
	return Map(out, func(k K) K { return k })
}

// ConvertWithErr returns a new slice with elements transformed by the given function with another type.
func ConvertWithErr[T, K any](input []T, transform func(T) (K, error)) ([]K, error) {
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
	out := make(map[K]T2, len(input))
	for k, v := range input {
		out[k] = transform(v)
	}
	return out
}

// ConvertMapWithErr returns a new map with elements transformed by the given function with another type.
func ConvertMapWithErr[K comparable, T1, T2 any](input map[K]T1, transform func(T1) (T2, error)) (map[K]T2, error) {
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

// FilterMap returns a new map with elements filtered by the given filter function.
func FilterMap[K comparable, T any](input map[K]T, filter func(K, T) bool) map[K]T {
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
	out := make([]T, len(input))
	copy(out, input)
	return out
}

// CopyMap returns a copy of a provided map.
func CopyMap[K comparable, T any](input map[K]T) map[K]T {
	out := make(map[K]T, len(input))
	for k, v := range input {
		out[k] = v
	}
	return out
}

// WithoutEmpty returns a new slice without empty elements.
func WithoutEmpty[T comparable](input []T) []T {
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
	out := make([]K, 0, len(input))
	for k := range input {
		out = append(out, k)
	}
	return out
}

// KeysIf returns a new slice with keys of a provided map filtered by the given filter function.
func KeysIf[K comparable, T any](input map[K]T, filter func(K, T) bool) []K {
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
	out := make([]T, 0, len(input))
	for _, v := range input {
		out = append(out, v)
	}
	return out
}

// ValuesIf returns a new slice with values of a provided map filtered by the given filter function.
func ValuesIf[K comparable, T any](input map[K]T, filter func(K, T) bool) []T {
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
	var empty T
	out := make(map[K]T, len(input))
	for k, v := range input {
		if v != empty {
			out[k] = v
		}
	}
	return out
}
