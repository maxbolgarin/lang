// Package lang provides useful generic oneliners to work with variables and pointers.
package lang

import "time"

// Ptr returns a pointer to a provided argument. It is useful to get an address of a literal.
//
//	// a := &"some literal" // won't compile
//	a := Ptr("some literal")
func Ptr[T any](v T) *T {
	return &v
}

// Check returns the first argument if it is not default, else returns the second one.
//
//	a := ""
//	b := "foo"
//	c := Check(a, b) // c == "foo"
func Check[T comparable](v1, v2 T) T {
	var empty T
	if v1 != empty {
		return v1
	}
	return v2
}

// CheckPtr returns dereference of the first argument (pointer) if it is not nil, else returns the second one.
//
//	a := ""
//	b := "foo"
//	c := CheckPtr(&a, b)  // c == ""
//	d := CheckPtr(nil, b) // d == "foo"
func CheckPtr[T any](v1 *T, v2 T) T {
	if v1 != nil {
		return *v1
	}
	return v2
}

// Deref returns dereference of the pointer if it is not nil, else returns the default value.
//
//	var a *int
//	b := 123
//	c := Deref(a) // c == 0
//	d := Deref(&b) // d == 123
func Deref[T any](v *T) T {
	if v == nil {
		var empty T
		return empty
	}
	return *v
}

// CheckTime returns the first time if it is not zero, second one elsewhere.
//
//	a := time.Time{}
//	b := time.Now()
//	c := CheckPtr(a, b)  // c == time.Now()
func CheckTime(v1 time.Time, v2 time.Time) time.Time {
	if !v1.IsZero() {
		return v1
	}
	return v2
}

// CheckIndex returns the value and true if the index is not out of bounds.
//
//	a := []int{1, 2, 3}
//	b, ok := CheckIndex(a, 2) // b == 3 && ok == true
//	c, ok := CheckIndex(a, 4) // c == 0 && ok == false
func CheckIndex[T any](s []T, index int) (T, bool) {
	if len(s) <= index {
		var empty T
		return empty, false
	}
	return s[index], true
}

// Index returns the value if the index is not out of bounds.
//
//	a := []int{1, 2, 3}
//	b := Index(a, 2) // b == 3
//	c := Index(a, 4) // c == 0
func Index[T any](s []T, index int) T {
	out, _ := CheckIndex(s, index)
	return out
}

// First returns the first element of the slice if it is not empty.
//
//	var a []int
//	b := []string{"foo", "bar"}
//	c := First(a)  // c == 0
//	d := First(b)  // d == "foo"
func First[T any](s []T) T {
	return Index(s, 0)
}

// If returns ifTrue if condition is true, otherwise it returns ifFalse.
//
//	a := If(true, 1, 2)  // a == 1
//	b := If(false, 1, 2) // b == 2
func If[T any](cond bool, ifTrue, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}

// IfF executes the function if the condition is true.
//
// IfF(true, func() { println("foo") })  // foo
// IfF(false, func() { println("foo") }) // nothing
func IfF(cond bool, f func()) {
	if cond {
		f()
	}
}

// IfV executes the function if the value is not zero.
//
//	a := IfV(1, func() { println("foo") })  // foo
//	b := IfV(0, func() { println("foo") })  // nothing
func IfV[T comparable](v T, f func()) {
	var zero T
	if v != zero {
		f()
	}
}

// GetWithSep returns the value (first argument) with the separator (second argument),
// if the separator does not exist in the last index of the value.
//
//	a := GetWithSep("config", '/')  // a == "config/"
//	b := GetWithSep("config/", '/') // b == "config/"
//	c := GetWithSep("config/files", '/') // b == "config/files/"
func GetWithSep(value string, sep byte) string {
	if value == "" {
		return ""
	}
	if value[len(value)-1] == sep {
		return value
	}
	return value + string(sep)
}

// CheckSlice returns the first argument if it is not empty, else returns the second one.
//
//	a := []int{}
//	b := []string{"foo", "bar"}
//	c := CheckSlice(a, b)  // c == []string{"foo", "bar"}
func CheckSlice[T any](v1, v2 []T) []T {
	if len(v1) == 0 {
		return v2
	}
	return v1
}

// CheckSliceSingle returns the first argument if it is not empty, else returns the second one wrapped in a slice.
//
//	a := nil
//	b := "foo"
//	c := CheckSliceSingle(a, b)  // c == []string{"foo"}
func CheckSliceSingle[T any](s []T, v T) []T {
	if len(s) > 0 {
		return s
	}
	return []T{v}
}

// CheckMap returns the first argument if it is not empty, else returns the second one.
//
//	a := map[string]int{}
//	b := map[string]string{"foo": "bar"}
//	c := CheckMap(a, b)  // c == map[string]string{"foo": "bar"}
func CheckMap[K comparable, V any](v1, v2 map[K]V) map[K]V {
	if len(v1) == 0 {
		return v2
	}
	return v1
}

// CheckMapSingle returns the first argument if it is not empty, else returns the second one wrapped in a map.
//
//	a := nil
//	b := "foo"
//	c := CheckMapSingle(a, b)  // c == map[string]string{"foo": "bar"}
func CheckMapSingle[K comparable, V any](m map[K]V, k K, v V) map[K]V {
	if len(m) > 0 {
		return m
	}
	m[k] = v
	return m
}

// IsFound returns if the value is in the slice.
//
//	a := []int{1, 2, 3}
//	b := IsFound(a, 3)  // b == true
//	c := IsFound(a, 4)  // c == false
func IsFound[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// MaxLen returns the slice with the maximum length.
//
//	a := []int{1, 2, 3}
//	b := MaxLen(a, 2) // b == [1, 2]
func MaxLen[T any](s []T, max int) []T {
	if len(s) <= max {
		return s
	}
	return s[:max]
}

// AppendIfAll appends the value to the slice if it is not empty.
// All values must be different from zero to be appended.
//
//	b := []string{"foo", "bar"}
//	c := AppendIfAll(b, "foo")  // c == []string{"foo", "bar", "foo"}
//	d := AppendIfAll(b, "")     // d == []string{"foo", "bar"}
//	e := AppendIfAll(b, "foo", "")  // e == []string{"foo", "bar"}
//	f := AppendIfAll(b, "foo", "bar")  // f == []string{"foo", "bar", "foo", "bar"}
func AppendIfAll[T comparable](s []T, v ...T) []T {
	if len(v) == 0 {
		return s
	}
	var zero T
	for _, e := range v {
		if e == zero {
			return s
		}
	}
	return append(s, v...)
}

// AppendIfAny appends the value to the slice if it is not empty.
// Any value must be different from zero to be appended.
//
//	b := []string{"foo", "bar"}
//	c := AppendIfAny(b, "foo")  // c == []string{"foo", "bar", "foo"}
//	d := AppendIfAny(b, "")     // d == []string{"foo", "bar"}
//	e := AppendIfAny(b, "foo", "")  // e == []string{"foo", "bar", "foo"}
//	f := AppendIfAny(b, "foo", "bar")  // f == []string{"foo", "bar", "foo", "bar"}
func AppendIfAny[T comparable](s []T, v ...T) []T {
	if len(v) == 0 {
		return s
	}
	var zero T
	for _, e := range v {
		if e != zero {
			s = append(s, e)
		}
	}
	return s
}
