// Package lang provides useful generic oneliners to work with variables and pointers.
package lang

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

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

// CheckPtrs returns the first argument if it is not nil, else returns the second one.
//
//	a := ""
//	b := "foo"
//	c := CheckPtrs(&a, &b)  // c == &a
//	d := CheckPtrs(nil, &b) // d == &b
func CheckPtrs[T any](v1 *T, v2 *T) *T {
	if v1 != nil {
		return v1
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
//	c := CheckTime(a, b)  // c == time.Now()
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
	if index < 0 || len(s) <= index {
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
	if cond && f != nil {
		f()
	}
}

// IfV executes the function if the value is not zero.
//
//	a := IfV(1, func() { println("foo") })  // foo
//	b := IfV(0, func() { println("foo") })  // nothing
func IfV[T comparable](v T, f func()) {
	var zero T
	if v != zero && f != nil {
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
	// Initialize map if nil to prevent panic
	if m == nil {
		m = make(map[K]V)
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
func MaxLen[T any](s []T, targetMax int) []T {
	if targetMax < 0 {
		targetMax = 0
	}
	if len(s) <= targetMax {
		return s
	}
	return s[:targetMax]
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
	if s == nil {
		s = []T{}
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
	if s == nil {
		s = []T{}
	}
	var zero T
	for _, e := range v {
		if e != zero {
			s = append(s, e)
		}
	}
	return s
}

// ConvertValue converts a value to a constant type.
//
//	a := 1
//	b := ConvertValue(a, func(a int) string { return strconv.Itoa(a) }) // b == "1"
func ConvertValue[T, K any](v T, f func(T) K) K {
	if f == nil {
		var zero K
		return zero
	}
	return f(v)
}

// WrapError adds a context message to an error.
//
//	err := SomeFunction()
//	if err != nil {
//	    return Wrap(err, "failed to execute SomeFunction")
//	}
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// JoinErrors combines multiple errors into a single error.
//
//	err1 := SomeFunction1()
//	err2 := SomeFunction2()
//	if err := JoinErrors(err1, err2); err != nil {
//	    return err
//	}
func JoinErrors(errs ...error) error {
	var nonNilErrs []string
	for _, err := range errs {
		if err != nil {
			nonNilErrs = append(nonNilErrs, err.Error())
		}
	}
	if len(nonNilErrs) == 0 {
		return nil
	}
	return errors.New(strings.Join(nonNilErrs, "; "))
}

// TruncateString truncates a string to a maximum length and adds an ellipsis if necessary.
//
//	s := "Hello, world!"
//	t := TruncateString(s, 5, "...") // t == "Hello..."
func TruncateString(s string, maxLen int, ellipsis ...string) string {
	if maxLen <= 0 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	if len(ellipsis) > 0 {
		return s[:maxLen] + ellipsis[0]
	}
	return s[:maxLen]
}

// String returns the string representation of the value with the optional maximum length.
//
//	a := String("Hello") // a == "Hello"
//	b := String(123) // b == "123"
//	c := String(123.456) // c == "123.456"
//	d := String(true) // d == "true"
//	e := String(time.Now()) // e == "2021-01-01T00:00:00Z"
//	f := String([]byte("Hello, world!")) // f == "Hello, world!"
//	g := String([]byte("Hello, world!"), 5) // g == "Hello"
//	h := String(nil, 5) // h == ""
//	i := String(nil, 0) // i == ""
//	j := String(nil, -1) // j == ""
func String(s any, maxLenRaw ...int) string {
	if s == nil {
		return ""
	}

	var maxLen int
	if len(maxLenRaw) > 0 {
		maxLen = maxLenRaw[0]
		if maxLen <= 0 {
			return ""
		}
	}

	switch v := s.(type) {
	case string:
		res := v
		return TruncateString(res, Check(maxLen, len(res)))

	case []byte:
		res := string(v)
		return TruncateString(res, Check(maxLen, len(res)))

	case []rune:
		res := string(v)
		return TruncateString(res, Check(maxLen, len(res)))

	case time.Time:
		res := v.Format(time.RFC3339)
		return TruncateString(res, Check(maxLen, len(res)))

	case fmt.Stringer:
		res := v.String()
		return TruncateString(res, Check(maxLen, len(res)))

	case error:
		res := v.Error()
		return TruncateString(res, Check(maxLen, len(res)))

	case int:
		res := strconv.FormatInt(int64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case int8:
		res := strconv.FormatInt(int64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case int16:
		res := strconv.FormatInt(int64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case int32:
		res := strconv.FormatInt(int64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case int64:
		res := strconv.FormatInt(v, 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case uint:
		res := strconv.FormatUint(uint64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case uint8:
		res := strconv.FormatUint(uint64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case uint16:
		res := strconv.FormatUint(uint64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case uint32:
		res := strconv.FormatUint(uint64(v), 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case uint64:
		res := strconv.FormatUint(v, 10)
		return TruncateString(res, Check(maxLen, len(res)))

	case float32:
		res := strconv.FormatFloat(float64(v), 'f', -1, 32)
		return TruncateString(res, Check(maxLen, len(res)))

	case float64:
		res := strconv.FormatFloat(v, 'f', -1, 64)
		return TruncateString(res, Check(maxLen, len(res)))

	case bool:
		res := strconv.FormatBool(v)
		return TruncateString(res, Check(maxLen, len(res)))

	default:
		res := fmt.Sprintf("%v", s)
		return TruncateString(res, Check(maxLen, len(res)))
	}
}

// S is a shortcut for [String].
func S(s any, maxLenRaw ...int) string {
	return String(s, maxLenRaw...)
}

// Type returns the value of the target type if the value is not nil.
//
//	a := Type[string](123) // a == ""
//	b := Type[string](nil) // b == ""
//	c := Type[string]("foo") // c == "foo"
//	d := Type[someEnum]("foo") // d == "" (type someEnum string) !!!
//	var v any = someEnum("foo")
//	e := Type[someEnum](v) // e == "foo" e.Type() == someEnum
func Type[Target any](s any) Target {
	var zero Target
	if s == nil {
		return zero
	}
	v, ok := s.(Target)
	if !ok {
		return zero
	}
	return v
}

// Retry attempts to execute a function until it succeeds or reaches max attempts.
//
//	result, err := Retry(3, func() (string, error) {
//	    return CallExternalAPI()
//	})
func Retry[T any](maxAttempts int, f func() (T, error)) (T, error) {
	var lastErr error
	for i := 0; i < maxAttempts; i++ {
		result, err := f()
		if err == nil {
			return result, nil
		}
		lastErr = err
	}
	var zero T
	return zero, fmt.Errorf("failed after %d attempts: %w", maxAttempts, lastErr)
}

var ErrTimeout = errors.New("operation timed out")

// RunWithTimeout runs a function with a timeout.
//
//	result, err := RunWithTimeout(time.Second, func() (string, error) {
//	    return SlowOperation()
//	})
func RunWithTimeout[T any](timeout time.Duration, f func() (T, error)) (T, error) {
	var result T
	var err error
	var wg sync.WaitGroup

	done := make(chan struct{})
	wg.Add(1)

	go func() {
		defer wg.Done()
		result, err = f()
		close(done)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(timeout):
		var zero T
		return zero, ErrTimeout
	}
}
