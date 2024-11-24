package lang_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/maxbolgarin/lang"
)

func TestSliceToMap(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	expectedResult := map[string]int{
		"1": 10,
		"2": 20,
		"3": 30,
		"4": 40,
		"5": 50,
	}
	result := lang.SliceToMap(inputSlice, func(i int) (string, int) {
		return strconv.Itoa(i), i * 10
	})
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}
}

func TestSliceToMapByKey(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	expectedResult := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}
	result := lang.SliceToMapByKey(inputSlice, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}
}

func TestPairsToMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := map[int]int{1: 2, 3: 4, 5: 6}
	result := lang.PairsToMap(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	input = []int{1, 2, 3, 4, 5}
	expected = map[int]int{1: 2, 3: 4}
	result = lang.PairsToMap(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	input = []int{1}
	expected = map[int]int{}
	result = lang.PairsToMap(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestFilter(t *testing.T) {
	var input []int
	filterFunc := func(n int) bool {
		return n > 0
	}
	expected := []int{}
	result := lang.Filter(input, filterFunc)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	input = []int{1, 2, 3, 4, -1, -2, -3}
	result = lang.Filter(input, filterFunc)
	if !reflect.DeepEqual(input[:4], result) {
		t.Fatalf("Expected %v but got %v", input[:4], result)
	}
}

func TestMap(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	expectedResult := []int{10, 20, 30, 40, 50}
	result := lang.Map(inputSlice, func(i int) int {
		return i * 10
	})
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}
}

func TestConvert(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	expectedResult := []int64{10, 20, 30, 40, 50}
	result := lang.Convert(inputSlice, func(i int) int64 {
		return int64(i * 10)
	})
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}
}

func TestConvertWithErr(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	expectedResult := []int64{10, 20, 30, 40, 50}
	result, err := lang.ConvertWithErr(inputSlice, func(i int) (int64, error) {
		return int64(i * 10), nil
	})
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}

	_, err = lang.ConvertWithErr(inputSlice, func(i int) (int64, error) {
		return int64(i * 10), errors.New("some error")
	})
	if err == nil {
		t.Fatalf("Expected error but got %v", err)
	}
}

func TestConvertMap(t *testing.T) {
	inputMap := map[string]int{"a": 1, "b": 2, "c": 3}
	expectedResult := map[string]int64{"a": 10, "b": 20, "c": 30}
	result := lang.ConvertMap(inputMap, func(i int) int64 {
		return int64(i * 10)
	})
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}
}

func TestConvertMapWithErr(t *testing.T) {
	inputMap := map[string]int{"a": 1, "b": 2, "c": 3}
	expectedResult := map[string]int64{"a": 10, "b": 20, "c": 30}
	result, err := lang.ConvertMapWithErr(inputMap, func(i int) (int64, error) {
		return int64(i * 10), nil
	})
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Expected %v but got %v", expectedResult, result)
	}

	_, err = lang.ConvertMapWithErr(inputMap, func(i int) (int64, error) {
		return int64(i * 10), errors.New("some error")
	})
	if err == nil {
		t.Fatalf("Expected error but got %v", err)
	}
}

func TestCopy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := lang.Copy(input)
	if !reflect.DeepEqual(input, result) {
		t.Fatalf("Expected %v but got %v", input, result)
	}
	if &input[0] == &result[0] {
		t.Fatalf("Expected different pointers but got the same")
	}
}

func TestCopyMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	result := lang.CopyMap(input)
	if !reflect.DeepEqual(input, result) {
		t.Fatalf("Expected %v but got %v", input, result)
	}
}

func TestKeys(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []string{"a", "b", "c"}
	result := lang.Keys(input)
	sort.Sort(sort.StringSlice(result))
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestKeysIf(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []string{"b"}
	result := lang.KeysIf(input, func(k string, v int) bool {
		return v%2 == 0
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestValues(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []int{1, 2, 3}
	result := lang.Values(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestValuesIf(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []int{2}
	result := lang.ValuesIf(input, func(k string, v int) bool {
		return v%2 == 0
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestWithoutEmpty(t *testing.T) {
	input := []string{"foo", "", "bar"}
	expected := []string{"foo", "bar"}
	result := lang.WithoutEmpty(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestWithoutEmptyValues(t *testing.T) {
	input := map[string]string{"foo": "", "bar": "bar"}
	expected := map[string]string{"bar": "bar"}
	result := lang.WithoutEmptyValues(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestWithoutEmptyKeys(t *testing.T) {
	input := map[string]string{"": "aaa", "bar": "bar"}
	expected := map[string]string{"bar": "bar"}
	result := lang.WithoutEmptyKeys(input)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}
