package lang_test

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/maxbolgarin/lang"
)

func TestPtr(t *testing.T) {
	if v := *lang.Ptr("foo"); v != "foo" {
		t.Errorf("expected %q but got %q", "foo", v)
	}
	if v := *lang.Ptr(""); v != "" {
		t.Errorf("expected %q but got %q", "", v)
	}
}

func TestCheck(t *testing.T) {
	if v := lang.Check("foo", "bar"); v != "foo" {
		t.Errorf("expected %q but got %q", "foo", v)
	}
	if v := lang.Check("", "bar"); v != "bar" {
		t.Errorf("expected %q but got %q", "bar", v)
	}
	if v := lang.Check("foo", ""); v != "foo" {
		t.Errorf("expected %q but got %q", "foo", v)
	}
	if v := lang.Check("", ""); v != "" {
		t.Errorf("expected %q but got %q", "", v)
	}
}

func TestCheckPtr(t *testing.T) {
	a := "foo"
	if v := lang.CheckPtr(&a, "bar"); v != "foo" {
		t.Errorf("expected %q but got %q", "foo", v)
	}
	if v := lang.CheckPtr(nil, "bar"); v != "bar" {
		t.Errorf("expected %q but got %q", "bar", v)
	}
	if v := lang.CheckPtr(nil, ""); v != "" {
		t.Errorf("expected %q but got %q", "", v)
	}

	b := ""
	if v := lang.CheckPtr(&b, "bar"); v != "" {
		t.Errorf("expected %q but got %q", "", v)
	}
}

func TestDeref(t *testing.T) {
	a := 123
	if v := lang.Deref[int](nil); v != 0 {
		t.Errorf("expected %d but got %d", 0, v)
	}
	if v := lang.Deref(&a); v != 123 {
		t.Errorf("expected %d but got %d", 123, v)
	}
}

func TestCheckTime(t *testing.T) {
	a := time.Time{}
	b := time.Now()
	if v := lang.CheckTime(a, b); !v.Equal(b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckTime(b, a); !v.Equal(b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckTime(a, a); !v.IsZero() {
		t.Errorf("expected %v but got %v", a, v)
	}
	if v := lang.CheckTime(b, b); !v.Equal(b) {
		t.Errorf("expected %v but got %v", b, v)
	}
}

func TestFirst(t *testing.T) {
	t.Run("EmptySlice", func(t *testing.T) {
		var a []int
		result := lang.First(a)
		if result != 0 {
			t.Errorf("expected %d but got %d", 0, result)
		}
	})

	t.Run("EmptySlice2", func(t *testing.T) {
		var b []string
		result := lang.First(b)
		if result != "" {
			t.Errorf("expected %q but got %q", "", result)
		}
	})

	t.Run("NotEmptySlice", func(t *testing.T) {
		b := []string{"foo", "bar"}
		result := lang.First(b)
		if result != "foo" {
			t.Errorf("expected %q but got %q", "foo", result)
		}
	})
}

func TestCheckIndex(t *testing.T) {
	t.Run("EmptySlice", func(t *testing.T) {
		var a []int
		out, ok := lang.CheckIndex(a, 0)
		if out != 0 || ok {
			t.Errorf("expected %d but got %d and ok:%v", 0, out, ok)
		}
	})

	t.Run("EmptySlice2", func(t *testing.T) {
		var b []string
		out, ok := lang.CheckIndex(b, 0)
		if out != "" || ok {
			t.Errorf("expected %q but got %q and ok:%v", "", out, ok)
		}
	})

	t.Run("NotEmptySlice", func(t *testing.T) {
		b := []string{"foo", "bar"}
		out, ok := lang.CheckIndex(b, 1)
		if out != "bar" || !ok {
			t.Errorf("expected %q but got %q and ok:%v", "bar", out, ok)
		}
	})

	t.Run("NotEmptySliceWrongIndex", func(t *testing.T) {
		b := []string{"foo", "bar"}
		out, ok := lang.CheckIndex(b, 2)
		if out != "" || ok {
			t.Errorf("expected %q but got %q and ok:%v", "", out, ok)
		}
	})
}

func TestIndex(t *testing.T) {
	t.Run("EmptySlice", func(t *testing.T) {
		var a []int
		out := lang.Index(a, 0)
		if out != 0 {
			t.Errorf("expected %d but got %d", 0, out)
		}
	})

	t.Run("EmptySlice2", func(t *testing.T) {
		var b []string
		out := lang.Index(b, 0)
		if out != "" {
			t.Errorf("expected %q but got %q", "", out)
		}
	})

	t.Run("NotEmptySlice", func(t *testing.T) {
		b := []string{"foo", "bar"}
		out := lang.Index(b, 1)
		if out != "bar" {
			t.Errorf("expected %q but got %q", "bar", out)
		}
	})

	t.Run("NotEmptySliceWrongIndex", func(t *testing.T) {
		b := []string{"foo", "bar"}
		out := lang.Index(b, 2)
		if out != "" {
			t.Errorf("expected %q but got %q", "", out)
		}
	})
}

func TestGetWithSep(t *testing.T) {
	testCases := []struct {
		value string
		sep   byte
		want  string
	}{
		{
			"config",
			'/',
			"config/",
		},
		{
			"config/",
			'/',
			"config/",
		},
		{
			"config/files",
			'/',
			"config/files/",
		},
		{
			"",
			'/',
			"",
		},
	}

	for _, tc := range testCases {
		if v := lang.GetWithSep(tc.value, tc.sep); v != tc.want {
			t.Errorf("expected %q but got %q", tc.want, v)
		}
	}
}

func TestIf(t *testing.T) {
	if v := lang.If(true, "foo", "bar"); v != "foo" {
		t.Errorf("expected %q but got %q", "foo", v)
	}
	if v := lang.If(false, "foo", "bar"); v != "bar" {
		t.Errorf("expected %q but got %q", "bar", v)
	}
}

func TestIfF(t *testing.T) {
	var a string
	lang.IfF(true, func() { a = "foo" })
	if a != "foo" {
		t.Errorf("expected %q but got %q", "foo", a)
	}

	var b string
	lang.IfF(false, func() { b = "foo" })
	if b != "" {
		t.Errorf("expected %q but got %q", "", b)
	}
}

func TestIfV(t *testing.T) {
	var a string
	lang.IfV("0", func() { a = "foo" })
	if a != "foo" {
		t.Errorf("expected %q but got %q", "foo", a)
	}

	var b string
	lang.IfV("", func() { b = "foo" })
	if b != "" {
		t.Errorf("expected %q but got %q", "", b)
	}

	var aa string
	ptr := &a
	lang.IfV(ptr, func() { aa = "foo" })
	if aa != "foo" {
		t.Errorf("expected %q but got %q", "foo", aa)
	}

	var bb string
	ptr = nil
	lang.IfV(ptr, func() { bb = "foo" })
	if bb != "" {
		t.Errorf("expected %q but got %q", "", bb)
	}
}

func TestCheckSlice(t *testing.T) {
	a := []string{}
	b := []string{"foo", "bar"}
	if v := lang.CheckSlice(a, b); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckSlice(b, a); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckSlice(a, a); !reflect.DeepEqual(v, a) {
		t.Errorf("expected %v but got %v", a, v)
	}
	if v := lang.CheckSlice(b, b); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
}

func TestCheckSliceSingle(t *testing.T) {
	a := []string{}
	b := []string{"foo", "bar"}
	c := []string{"foo"}
	if v := lang.CheckSliceSingle(a, "foo"); !reflect.DeepEqual(v, c) {
		t.Errorf("expected %v but got %v", c, v)
	}
	if v := lang.CheckSliceSingle(b, "foo"); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
}

func TestCheckMap(t *testing.T) {
	a := map[string]string{}
	b := map[string]string{"foo": "bar"}
	if v := lang.CheckMap(a, b); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckMap(b, a); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
	if v := lang.CheckMap(a, a); !reflect.DeepEqual(v, a) {
		t.Errorf("expected %v but got %v", a, v)
	}
	if v := lang.CheckMap(b, b); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
}

func TestCheckMapSingle(t *testing.T) {
	a := map[string]string{}
	b := map[string]string{"foo": "bar"}
	c := map[string]string{"foo2": "bar2"}
	if v := lang.CheckMapSingle(a, "foo2", "bar2"); !reflect.DeepEqual(v, c) {
		t.Errorf("expected %v but got %v", c, v)
	}
	if v := lang.CheckMapSingle(b, "foo2", "bar2"); !reflect.DeepEqual(v, b) {
		t.Errorf("expected %v but got %v", b, v)
	}
}

func TestIsFound(t *testing.T) {
	a := []string{"foo", "bar"}
	if v := lang.IsFound(a, "foo"); !v {
		t.Errorf("expected %v but got %v", true, v)
	}
	if v := lang.IsFound(a, "bar"); !v {
		t.Errorf("expected %v but got %v", true, v)
	}
	if v := lang.IsFound(a, "baz"); v {
		t.Errorf("expected %v but got %v", false, v)
	}
}

func TestMaxLen(t *testing.T) {
	a := []string{}
	v := lang.MaxLen(a, 2)
	if len(v) != 0 {
		t.Errorf("expected %d but got %d", 0, len(v))
	}
	a = []string{"foo", "bar", "baz"}
	v = lang.MaxLen(a, 2)
	if len(v) != 2 {
		t.Errorf("expected %d but got %d", 2, len(v))
	}
	for i := range v {
		if v[i] != a[i] {
			t.Errorf("expected %q but got %q", a[i], v[i])
		}
	}
}

func TestAppendIfAll(t *testing.T) {
	a := []string{}
	if v := lang.AppendIfAll(a, "foo"); !reflect.DeepEqual(v, []string{"foo"}) {
		t.Errorf("expected %v but got %v", []string{"foo"}, v)
	}
	b := []string{"foo", "bar"}
	c := lang.AppendIfAll(b, "foo")
	if !reflect.DeepEqual(c, []string{"foo", "bar", "foo"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar", "foo"}, c)
	}
	d := lang.AppendIfAll(b, "")
	if !reflect.DeepEqual(d, b) {
		t.Errorf("expected %v but got %v", b, d)
	}
	e := lang.AppendIfAll(b, "foo", "")
	if !reflect.DeepEqual(e, []string{"foo", "bar"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar"}, e)
	}
	f := lang.AppendIfAll(b, "foo", "bar")
	if !reflect.DeepEqual(f, []string{"foo", "bar", "foo", "bar"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar", "foo", "bar"}, f)
	}
	f = lang.AppendIfAll(b)
	if !reflect.DeepEqual(f, b) {
		t.Errorf("expected %v but got %v", b, f)
	}
}

func TestAppendIfAny(t *testing.T) {
	a := []string{}
	if v := lang.AppendIfAny(a, "foo"); !reflect.DeepEqual(v, []string{"foo"}) {
		t.Errorf("expected %v but got %v", []string{"foo"}, v)
	}
	b := []string{"foo", "bar"}
	c := lang.AppendIfAny(b, "foo")
	if !reflect.DeepEqual(c, []string{"foo", "bar", "foo"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar", "foo"}, c)
	}
	d := lang.AppendIfAny(b, "")
	if !reflect.DeepEqual(d, []string{"foo", "bar"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar"}, d)
	}
	e := lang.AppendIfAny(b, "foo", "")
	if !reflect.DeepEqual(e, []string{"foo", "bar", "foo"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar", "foo"}, e)
	}
	f := lang.AppendIfAny(b, "foo", "bar")
	if !reflect.DeepEqual(f, []string{"foo", "bar", "foo", "bar"}) {
		t.Errorf("expected %v but got %v", []string{"foo", "bar", "foo", "bar"}, f)
	}
	f = lang.AppendIfAny(b)
	if !reflect.DeepEqual(f, b) {
		t.Errorf("expected %v but got %v", b, f)
	}
}

func TestConvertValue(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		input := 42
		expected := "42"
		result := lang.ConvertValue(input, strconv.Itoa)
		if result != expected {
			t.Errorf("ConvertValue(%v, strconv.Itoa) = %v, want %v", input, result, expected)
		}
	})

	t.Run("string to int", func(t *testing.T) {
		input := "42"
		expected := 42
		result := lang.ConvertValue(input, func(s string) int {
			v, _ := strconv.Atoi(s)
			return v
		})
		if result != expected {
			t.Errorf("ConvertValue(%v, func) = %v, want %v", input, result, expected)
		}
	})

	t.Run("struct to map", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		input := Person{Name: "Alice", Age: 30}
		expected := map[string]interface{}{"name": "Alice", "age": 30}

		result := lang.ConvertValue(input, func(p Person) map[string]interface{} {
			return map[string]interface{}{
				"name": p.Name,
				"age":  p.Age,
			}
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ConvertValue(%v, func) = %v, want %v", input, result, expected)
		}
	})

	t.Run("identity function", func(t *testing.T) {
		input := "unchanged"
		expected := "unchanged"
		result := lang.ConvertValue(input, func(s string) string { return s })
		if result != expected {
			t.Errorf("ConvertValue(%v, func) = %v, want %v", input, result, expected)
		}
	})
}

func TestCheckIndex_EdgeCases(t *testing.T) {
	t.Run("NegativeIndex", func(t *testing.T) {
		s := []string{"foo", "bar"}
		val, ok := lang.CheckIndex(s, -1)
		if ok || val != "" {
			t.Errorf("Expected (empty, false) for negative index, got (%v, %v)", val, ok)
		}
	})
}

func TestIfF_NilFunction(t *testing.T) {
	// This test just verifies no panic occurs
	lang.IfF(true, nil)
	lang.IfF(false, nil)
}

func TestIfV_NilFunction(t *testing.T) {
	// This test just verifies no panic occurs
	val := "test"
	lang.IfV(val, nil)

	var empty string
	lang.IfV(empty, nil)
}

func TestCheckMapSingle_NilMap(t *testing.T) {
	var nilMap map[string]int
	result := lang.CheckMapSingle(nilMap, "key", 42)

	if result == nil {
		t.Error("Expected initialized map, got nil")
	}

	if val, exists := result["key"]; !exists || val != 42 {
		t.Errorf("Expected map with {'key': 42}, got %v", result)
	}
}

func TestMaxLen_NegativeMax(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	result := lang.MaxLen(s, -5)

	if len(result) != 0 {
		t.Errorf("Expected empty slice for negative max, got %v", result)
	}
}

func TestAppendIfAll_NilSlice(t *testing.T) {
	var nilSlice []string
	result := lang.AppendIfAll(nilSlice, "foo", "bar")

	expected := []string{"foo", "bar"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Also test with a zero value in the arguments
	result = lang.AppendIfAll(nilSlice, "foo", "")
	if len(result) != 0 {
		t.Errorf("Expected empty slice when one argument is zero, got %v", result)
	}
}

func TestAppendIfAny_NilSlice(t *testing.T) {
	var nilSlice []string
	result := lang.AppendIfAny(nilSlice, "foo", "bar")

	expected := []string{"foo", "bar"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Also test with a zero value in the arguments
	result = lang.AppendIfAny(nilSlice, "foo", "")
	expected = []string{"foo"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConvertValue_NilFunction(t *testing.T) {
	var nilFunc func(int) string
	result := lang.ConvertValue(42, nilFunc)

	if result != "" {
		t.Errorf("Expected empty string for nil function, got %v", result)
	}
}

// Add test cases to verify fixes for pointer and nil handling

func TestPtr_Consistency(t *testing.T) {
	val := "test"
	ptr := lang.Ptr(val)

	if *ptr != val {
		t.Errorf("Expected pointer to %v, got pointer to %v", val, *ptr)
	}

	// Verify that modifying original doesn't affect pointer value
	val = "changed"
	if *ptr == val {
		t.Errorf("Expected pointer to keep original value, got %v", *ptr)
	}
}

func TestCheckPtr_EdgeCases(t *testing.T) {
	// Already covered in main tests, but add explicit nil test
	var nilPtr *string
	result := lang.CheckPtr(nilPtr, "default")
	if result != "default" {
		t.Errorf("Expected default value for nil pointer, got %v", result)
	}
}

func TestDeref_Consistency(t *testing.T) {
	// Already covered in main tests, but add explicit check
	val := 42
	ptr := &val
	result := lang.Deref(ptr)

	if result != val {
		t.Errorf("Expected %v, got %v", val, result)
	}

	// Check nil deref
	var nilPtr *int
	zeroResult := lang.Deref(nilPtr)
	if zeroResult != 0 {
		t.Errorf("Expected 0 for nil pointer deref, got %v", zeroResult)
	}
}

func TestIndex_NegativeIndex(t *testing.T) {
	s := []string{"foo", "bar"}
	result := lang.Index(s, -1)

	if result != "" {
		t.Errorf("Expected empty string for negative index, got %v", result)
	}
}
