package lang_test

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/maxbolgarin/lang"
)

func TestSliceToMap(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		transform func(int) (string, string)
		want      map[string]string
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{"1": "1", "2": "4", "3": "9"},
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
		{
			name:  "duplicate keys",
			input: []int{1, 2, 1},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{"1": "1", "2": "4"}, // Last value overwrites
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.SliceToMap(tt.input, tt.transform)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToMapByKey(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		key   func(int) string
		want  map[string]int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			key: func(i int) string {
				return strconv.Itoa(i)
			},
			want: map[string]int{"1": 1, "2": 2, "3": 3},
		},
		{
			name:  "empty slice",
			input: []int{},
			key: func(i int) string {
				return strconv.Itoa(i)
			},
			want: map[string]int{},
		},
		{
			name:  "nil slice",
			input: nil,
			key: func(i int) string {
				return strconv.Itoa(i)
			},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.SliceToMapByKey(tt.input, tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMapByKey() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Mapping(tt.input, tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMapByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairsToMap(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string]string
	}{
		{
			name:  "normal usage",
			input: []string{"a", "1", "b", "2", "c", "3"},
			want:  map[string]string{"a": "1", "b": "2", "c": "3"},
		},
		{
			name:  "empty slice",
			input: []string{},
			want:  map[string]string{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  map[string]string{},
		},
		{
			name:  "odd length",
			input: []string{"a", "1", "b", "2", "c"},
			want:  map[string]string{"a": "1", "b": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.PairsToMap(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairsToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		filter func(int) bool
		want   []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{2, 4},
		},
		{
			name:  "empty slice",
			input: []int{},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: nil,
		},
		{
			name:  "all match",
			input: []int{2, 4, 6},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{2, 4, 6},
		},
		{
			name:  "none match",
			input: []int{1, 3, 5},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Filter(tt.input, tt.filter)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		transform func(int) int
		want      []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) int {
				return i * 2
			},
			want: []int{2, 4, 6},
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) int {
				return i * 2
			},
			want: []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) int {
				return i * 2
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Map(tt.input, tt.transform)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		initial int
		f       func(int, int) int
		want    int
	}{
		{
			name:    "sum",
			input:   []int{1, 2, 3, 4},
			initial: 0,
			f: func(acc, val int) int {
				return acc + val
			},
			want: 10,
		},
		{
			name:    "product",
			input:   []int{1, 2, 3, 4},
			initial: 1,
			f: func(acc, val int) int {
				return acc * val
			},
			want: 24,
		},
		{
			name:    "empty slice",
			input:   []int{},
			initial: 100,
			f: func(acc, val int) int {
				return acc + val
			},
			want: 100,
		},
		{
			name:    "nil slice",
			input:   nil,
			initial: 100,
			f: func(acc, val int) int {
				return acc + val
			},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Reduce(tt.input, tt.initial, tt.f)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		transform func(int) string
		want      []string
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) string {
				return strconv.Itoa(i)
			},
			want: []string{"1", "2", "3"},
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) string {
				return strconv.Itoa(i)
			},
			want: []string{},
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) string {
				return strconv.Itoa(i)
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Convert(tt.input, tt.transform)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertWithErr(t *testing.T) {
	testErr := errors.New("test error")
	tests := []struct {
		name      string
		input     []int
		transform func(int) (string, error)
		want      []string
		wantErr   bool
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) (string, error) {
				return strconv.Itoa(i), nil
			},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
		{
			name:  "with error",
			input: []int{1, 2, 3},
			transform: func(i int) (string, error) {
				if i == 2 {
					return "", testErr
				}
				return strconv.Itoa(i), nil
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) (string, error) {
				return strconv.Itoa(i), nil
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) (string, error) {
				return strconv.Itoa(i), nil
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lang.ConvertWithErr(tt.input, tt.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertMap(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		transform func(int) string
		want      map[string]string
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(i int) string {
				return strconv.Itoa(i * i)
			},
			want: map[string]string{"a": "1", "b": "4", "c": "9"},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			transform: func(i int) string {
				return strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
		{
			name:  "nil map",
			input: nil,
			transform: func(i int) string {
				return strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ConvertMap(tt.input, tt.transform)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertMapWithErr(t *testing.T) {
	testErr := errors.New("test error")
	tests := []struct {
		name      string
		input     map[string]int
		transform func(int) (string, error)
		want      map[string]string
		wantErr   bool
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(i int) (string, error) {
				return strconv.Itoa(i * i), nil
			},
			want:    map[string]string{"a": "1", "b": "4", "c": "9"},
			wantErr: false,
		},
		{
			name:  "with error",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(i int) (string, error) {
				if i == 2 {
					return "", testErr
				}
				return strconv.Itoa(i * i), nil
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty map",
			input: map[string]int{},
			transform: func(i int) (string, error) {
				return strconv.Itoa(i * i), nil
			},
			want:    map[string]string{},
			wantErr: false,
		},
		{
			name:  "nil map",
			input: nil,
			transform: func(i int) (string, error) {
				return strconv.Itoa(i * i), nil
			},
			want:    map[string]string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lang.ConvertMapWithErr(tt.input, tt.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertMapWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertMapWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFromMap(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		transform func(string, int) string
		want      []string
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(k string, v int) string {
				return k + ":" + strconv.Itoa(v)
			},
			want: []string{"a:1", "b:2", "c:3"},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			transform: func(k string, v int) string {
				return k + ":" + strconv.Itoa(v)
			},
			want: []string{},
		},
		{
			name:  "nil map",
			input: nil,
			transform: func(k string, v int) string {
				return k + ":" + strconv.Itoa(v)
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ConvertFromMap(tt.input, tt.transform)

			// Sort for deterministic comparison since map iteration order is non-deterministic
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFromMapWithErr(t *testing.T) {
	testErr := errors.New("test error")
	tests := []struct {
		name      string
		input     map[string]int
		transform func(string, int) (string, error)
		want      []string
		wantErr   bool
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(k string, v int) (string, error) {
				return k + ":" + strconv.Itoa(v), nil
			},
			want:    []string{"a:1", "b:2", "c:3"},
			wantErr: false,
		},
		{
			name:  "with error",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(k string, v int) (string, error) {
				if v == 2 {
					return "", testErr
				}
				return k + ":" + strconv.Itoa(v), nil
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty map",
			input: map[string]int{},
			transform: func(k string, v int) (string, error) {
				return k + ":" + strconv.Itoa(v), nil
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name:  "nil map",
			input: nil,
			transform: func(k string, v int) (string, error) {
				return k + ":" + strconv.Itoa(v), nil
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lang.ConvertFromMapWithErr(tt.input, tt.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFromMapWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Sort for deterministic comparison since map iteration order is non-deterministic
				sort.Strings(got)
				sort.Strings(tt.want)

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ConvertFromMapWithErr() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestConvertToMap(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		transform func(int) (string, string)
		want      map[string]string
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{"1": "1", "2": "4", "3": "9"},
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{},
		},
		{
			name:  "duplicate keys",
			input: []int{1, 2, 1},
			transform: func(i int) (string, string) {
				return strconv.Itoa(i), strconv.Itoa(i * i)
			},
			want: map[string]string{"1": "1", "2": "4"}, // Last value wins
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ConvertToMap(tt.input, tt.transform)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToMapWithErr(t *testing.T) {
	testErr := errors.New("test error")
	tests := []struct {
		name      string
		input     []int
		transform func(int) (string, string, error)
		want      map[string]string
		wantErr   bool
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			transform: func(i int) (string, string, error) {
				return strconv.Itoa(i), strconv.Itoa(i * i), nil
			},
			want:    map[string]string{"1": "1", "2": "4", "3": "9"},
			wantErr: false,
		},
		{
			name:  "with error",
			input: []int{1, 2, 3},
			transform: func(i int) (string, string, error) {
				if i == 2 {
					return "", "", testErr
				}
				return strconv.Itoa(i), strconv.Itoa(i * i), nil
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty slice",
			input: []int{},
			transform: func(i int) (string, string, error) {
				return strconv.Itoa(i), strconv.Itoa(i * i), nil
			},
			want:    map[string]string{},
			wantErr: false,
		},
		{
			name:  "nil slice",
			input: nil,
			transform: func(i int) (string, string, error) {
				return strconv.Itoa(i), strconv.Itoa(i * i), nil
			},
			want:    map[string]string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lang.ConvertToMapWithErr(tt.input, tt.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToMapWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToMapWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]int
		filter func(string, int) bool
		want   map[string]int
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: map[string]int{"b": 2, "d": 4},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: map[string]int{},
		},
		{
			name:  "nil map",
			input: nil,
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: map[string]int{},
		},
		{
			name:  "all match",
			input: map[string]int{"b": 2, "d": 4, "f": 6},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: map[string]int{"b": 2, "d": 4, "f": 6},
		},
		{
			name:  "none match",
			input: map[string]int{"a": 1, "c": 3, "e": 5},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.FilterMap(tt.input, tt.filter)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Copy(tt.input)

			// Verify it's a different slice (not the same reference)
			if len(tt.input) > 0 && len(got) > 0 {
				// Modify the copy to ensure it doesn't affect the original
				got[0] = got[0] + 100
				if tt.input[0] == got[0] {
					t.Errorf("Copy() returned the same slice reference, not a copy")
				}
			}

			// Reset for comparison
			if len(got) > 0 {
				got[0] = tt.want[0]
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			want:  map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			want:  map[string]int{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.CopyMap(tt.input)

			// Verify it's a different map (not the same reference)
			if len(tt.input) > 0 {
				// Add a new key to the copy
				got["test"] = 999
				if _, exists := tt.input["test"]; exists {
					t.Errorf("CopyMap() returned the same map reference, not a copy")
				}

				// Remove for comparison
				delete(got, "test")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  []string
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			want:  []string{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Keys(tt.input)

			// Sort for deterministic comparison since map iteration order is non-deterministic
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeysIf(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]int
		filter func(string, int) bool
		want   []string
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []string{"b", "d"},
		},
		{
			name:  "all match",
			input: map[string]int{"b": 2, "d": 4, "f": 6},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []string{"b", "d", "f"},
		},
		{
			name:  "none match",
			input: map[string]int{"a": 1, "c": 3, "e": 5},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []string{},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []string{},
		},
		{
			name:  "nil map",
			input: nil,
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.KeysIf(tt.input, tt.filter)

			// Sort for deterministic comparison since map iteration order is non-deterministic
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeysIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  []int
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			want:  []int{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Values(tt.input)

			// Sort for deterministic comparison since map iteration order is non-deterministic
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesIf(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]int
		filter func(string, int) bool
		want   []int
	}{
		{
			name:  "normal usage",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []int{2, 4},
		},
		{
			name:  "all match",
			input: map[string]int{"b": 2, "d": 4, "f": 6},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []int{2, 4, 6},
		},
		{
			name:  "none match",
			input: map[string]int{"a": 1, "c": 3, "e": 5},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name:  "nil map",
			input: nil,
			filter: func(k string, v int) bool {
				return v%2 == 0
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ValuesIf(tt.input, tt.filter)

			// Sort for deterministic comparison since map iteration order is non-deterministic
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithoutEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "normal usage",
			input: []string{"a", "", "b", "", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "no empty values",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "all empty values",
			input: []string{"", "", ""},
			want:  []string{},
		},
		{
			name:  "empty slice",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.WithoutEmpty(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithoutEmptyKeys(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int
	}{
		{
			name:  "normal usage",
			input: map[string]int{"": 1, "b": 2, "c": 3, "d": 4},
			want:  map[string]int{"b": 2, "c": 3, "d": 4},
		},
		{
			name:  "no empty keys",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			want:  map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:  "all empty keys",
			input: map[string]int{"": 1, "x": 2, "y": 3},
			want:  map[string]int{"x": 2, "y": 3},
		},
		{
			name:  "empty map",
			input: map[string]int{},
			want:  map[string]int{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.WithoutEmptyKeys(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutEmptyKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithoutEmptyValues(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
		want  map[string]string
	}{
		{
			name:  "normal usage",
			input: map[string]string{"a": "", "b": "value", "c": "", "d": "another"},
			want:  map[string]string{"b": "value", "d": "another"},
		},
		{
			name:  "no empty values",
			input: map[string]string{"a": "val1", "b": "val2", "c": "val3"},
			want:  map[string]string{"a": "val1", "b": "val2", "c": "val3"},
		},
		{
			name:  "all empty values",
			input: map[string]string{"a": "", "b": "", "c": ""},
			want:  map[string]string{},
		},
		{
			name:  "empty map",
			input: map[string]string{},
			want:  map[string]string{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.WithoutEmptyValues(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutEmptyValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "normal usage",
			input: []string{"a", "", "b", "", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "no empty values",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "all empty values",
			input: []string{"", "", ""},
			want:  []string{},
		},
		{
			name:  "empty slice",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.NotEmpty(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEmptyMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
		want  map[string]string
	}{
		{
			name:  "normal usage",
			input: map[string]string{"": "val1", "b": "", "c": "val3", "d": "", "e": "val5"},
			want:  map[string]string{"c": "val3", "e": "val5"},
		},
		{
			name:  "no empty keys or values",
			input: map[string]string{"a": "val1", "b": "val2", "c": "val3"},
			want:  map[string]string{"a": "val1", "b": "val2", "c": "val3"},
		},
		{
			name:  "all empty keys or values",
			input: map[string]string{"a": "", "c": ""},
			want:  map[string]string{},
		},
		{
			name:  "empty map",
			input: map[string]string{},
			want:  map[string]string{},
		},
		{
			name:  "nil map",
			input: nil,
			want:  map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.NotEmptyMap(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotEmptyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitByChunkSize(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		chunkSize int
		want      [][]int
	}{
		{
			name:      "normal usage",
			input:     []int{1, 2, 3, 4, 5, 6, 7},
			chunkSize: 3,
			want:      [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:      "chunk size equals length",
			input:     []int{1, 2, 3},
			chunkSize: 3,
			want:      [][]int{{1, 2, 3}},
		},
		{
			name:      "chunk size greater than length",
			input:     []int{1, 2, 3},
			chunkSize: 5,
			want:      [][]int{{1, 2, 3}},
		},
		{
			name:      "zero chunk size",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: 0,
			want:      [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:      "negative chunk size",
			input:     []int{1, 2, 3, 4, 5},
			chunkSize: -2,
			want:      [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:      "empty slice",
			input:     []int{},
			chunkSize: 3,
			want:      [][]int{},
		},
		{
			name:      "nil slice",
			input:     nil,
			chunkSize: 3,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.SplitByChunkSize(tt.input, tt.chunkSize)
			if !reflect.DeepEqual(got, tt.want) {
				if len(got) != len(tt.want) {
					t.Errorf("SplitByChunkSize() = %v, want %v", got, tt.want)
				}
				for i, chunk := range got {
					if !reflect.DeepEqual(chunk, tt.want[i]) {
						t.Errorf("SplitByChunkSize() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestFindFirst(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      int
		wantFound bool
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool {
				return i > 3
			},
			want:      4,
			wantFound: true,
		},
		{
			name:  "first element matches",
			input: []int{5, 2, 3, 4, 1},
			predicate: func(i int) bool {
				return i > 3
			},
			want:      5,
			wantFound: true,
		},
		{
			name:  "last element matches",
			input: []int{1, 2, 3, 2, 5},
			predicate: func(i int) bool {
				return i > 3
			},
			want:      5,
			wantFound: true,
		},
		{
			name:  "no match",
			input: []int{1, 2, 3},
			predicate: func(i int) bool {
				return i > 5
			},
			want:      0, // Zero value for int
			wantFound: false,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) bool {
				return i > 3
			},
			want:      0,
			wantFound: false,
		},
		{
			name:  "nil slice",
			input: nil,
			predicate: func(i int) bool {
				return i > 3
			},
			want:      0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := lang.FindFirst(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("FindFirst() got = %v, want %v", got, tt.want)
			}
			if found != tt.wantFound {
				t.Errorf("FindFirst() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

// Tests for newly added functions

func TestContains(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		element int
		want    bool
	}{
		{
			name:    "element exists",
			input:   []int{1, 2, 3, 4, 5},
			element: 3,
			want:    true,
		},
		{
			name:    "element does not exist",
			input:   []int{1, 2, 4, 5},
			element: 3,
			want:    false,
		},
		{
			name:    "empty slice",
			input:   []int{},
			element: 3,
			want:    false,
		},
		{
			name:    "nil slice",
			input:   nil,
			element: 3,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Contains(tt.input, tt.element)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:  "element exists",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool {
				return i > 3
			},
			want: true,
		},
		{
			name:  "element does not exist",
			input: []int{1, 2, 3},
			predicate: func(i int) bool {
				return i > 5
			},
			want: false,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) bool {
				return i > 3
			},
			want: false,
		},
		{
			name:  "nil slice",
			input: nil,
			predicate: func(i int) bool {
				return i > 3
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ContainsFunc(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		element string
		want    int
	}{
		{
			name:    "element exists",
			input:   []string{"a", "b", "c", "d"},
			element: "c",
			want:    2,
		},
		{
			name:    "element exists at start",
			input:   []string{"a", "b", "c"},
			element: "a",
			want:    0,
		},
		{
			name:    "element exists at end",
			input:   []string{"a", "b", "c"},
			element: "c",
			want:    2,
		},
		{
			name:    "element does not exist",
			input:   []string{"a", "b", "c"},
			element: "d",
			want:    -1,
		},
		{
			name:    "empty slice",
			input:   []string{},
			element: "a",
			want:    -1,
		},
		{
			name:    "nil slice",
			input:   nil,
			element: "a",
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.IndexOf(tt.input, tt.element)
			if got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		element string
		want    int
	}{
		{
			name:    "element exists once",
			input:   []string{"a", "b", "c", "d"},
			element: "c",
			want:    2,
		},
		{
			name:    "element exists multiple times",
			input:   []string{"a", "b", "c", "b", "d"},
			element: "b",
			want:    3,
		},
		{
			name:    "element exists at start",
			input:   []string{"a", "b", "c"},
			element: "a",
			want:    0,
		},
		{
			name:    "element exists at end",
			input:   []string{"a", "b", "c"},
			element: "c",
			want:    2,
		},
		{
			name:    "element does not exist",
			input:   []string{"a", "b", "c"},
			element: "d",
			want:    -1,
		},
		{
			name:    "empty slice",
			input:   []string{},
			element: "a",
			want:    -1,
		},
		{
			name:    "nil slice",
			input:   nil,
			element: "a",
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.LastIndexOf(tt.input, tt.element)
			if got != tt.want {
				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 2, 3, 1, 4, 5, 3},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "no duplicates",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "all duplicates",
			input: []int{1, 1, 1, 1},
			want:  []int{1},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Distinct(tt.input)

			// Compare as sets (order might be different but still correct)
			if len(got) != len(tt.want) {
				t.Errorf("Distinct() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Convert to map for easy lookup
			wantMap := make(map[int]bool)
			for _, v := range tt.want {
				wantMap[v] = true
			}

			for _, v := range got {
				if !wantMap[v] {
					t.Errorf("Distinct() result contains unexpected value %v", v)
				}
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "normal usage",
			a:    []int{1, 2, 3, 4},
			b:    []int{3, 4, 5, 6},
			want: []int{3, 4},
		},
		{
			name: "duplicates in inputs",
			a:    []int{1, 2, 3, 3, 4},
			b:    []int{3, 3, 4, 5, 6},
			want: []int{3, 4},
		},
		{
			name: "no common elements",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "one empty slice",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{},
		},
		{
			name: "both empty slices",
			a:    []int{},
			b:    []int{},
			want: []int{},
		},
		{
			name: "one nil slice",
			a:    []int{1, 2, 3},
			b:    nil,
			want: nil,
		},
		{
			name: "both nil slices",
			a:    nil,
			b:    nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Intersect(tt.a, tt.b)

			// Convert both to maps for set comparison
			gotMap := make(map[int]bool)
			for _, v := range got {
				gotMap[v] = true
			}

			wantMap := make(map[int]bool)
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if len(got) != len(tt.want) {
				t.Errorf("Intersect() length = %v, want %v", len(got), len(tt.want))
				return
			}

			for v := range gotMap {
				if !wantMap[v] {
					t.Errorf("Intersect() contains unexpected value %v", v)
				}
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{
			name:   "two slices",
			slices: [][]int{{1, 2, 3}, {3, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "three slices",
			slices: [][]int{{1, 2}, {2, 3}, {3, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "with duplicates",
			slices: [][]int{{1, 2, 2}, {2, 3, 3}, {3, 4, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "one empty slice",
			slices: [][]int{{1, 2, 3}, {}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "all empty slices",
			slices: [][]int{{}, {}},
			want:   []int{},
		},
		{
			name:   "one nil slice",
			slices: [][]int{{1, 2, 3}, nil},
			want:   []int{1, 2, 3},
		},
		{
			name:   "all nil slices",
			slices: [][]int{nil, nil},
			want:   nil,
		},
		{
			name:   "no slices",
			slices: [][]int{},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Union(tt.slices...)

			// Convert both to maps for set comparison
			gotMap := make(map[int]bool)
			for _, v := range got {
				gotMap[v] = true
			}

			wantMap := make(map[int]bool)
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if len(got) != len(tt.want) {
				t.Errorf("Union() length = %v, want %v", len(got), len(tt.want))
				return
			}

			for v := range gotMap {
				if !wantMap[v] {
					t.Errorf("Union() contains unexpected value %v", v)
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "normal usage",
			a:    []int{1, 2, 3, 4},
			b:    []int{3, 4, 5, 6},
			want: []int{1, 2},
		},
		{
			name: "no common elements",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{1, 2, 3},
		},
		{
			name: "all elements in common",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "with duplicates",
			a:    []int{1, 2, 2, 3, 3},
			b:    []int{2, 3},
			want: []int{1},
		},
		{
			name: "empty first slice",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "empty second slice",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "nil first slice",
			a:    nil,
			b:    []int{1, 2, 3},
			want: nil,
		},
		{
			name: "nil second slice",
			a:    []int{1, 2, 3},
			b:    nil,
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Difference(tt.a, tt.b)

			// Sort for consistent ordering
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Reverse(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "normal usage",
			input: [][]int{{1, 2}, {3, 4}, {5}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with empty slices",
			input: [][]int{{1, 2}, {}, {3, 4}, {}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "all empty slices",
			input: [][]int{{}, {}, {}},
			want:  []int{},
		},
		{
			name:  "empty outer slice",
			input: [][]int{},
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Flatten(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		size  int
		want  [][]int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			size:  3,
			want:  [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:  "size equals length",
			input: []int{1, 2, 3},
			size:  3,
			want:  [][]int{{1, 2, 3}},
		},
		{
			name:  "empty slice",
			input: []int{},
			size:  3,
			want:  [][]int{},
		},
		{
			name:  "nil slice",
			input: nil,
			size:  3,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Chunk(tt.input, tt.size)
			if !reflect.DeepEqual(got, tt.want) {
				if len(got) != len(tt.want) {
					t.Errorf("Chunk() = %v, want %v", got, tt.want)
				}
				for i, chunk := range got {
					if !reflect.DeepEqual(chunk, tt.want[i]) {
						t.Errorf("Chunk() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name  string
		input []Person
		keyFn func(Person) int
		want  map[int][]Person
	}{
		{
			name: "normal usage",
			input: []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 25},
				{"Dave", 30},
			},
			keyFn: func(p Person) int {
				return p.Age
			},
			want: map[int][]Person{
				25: {{"Alice", 25}, {"Charlie", 25}},
				30: {{"Bob", 30}, {"Dave", 30}},
			},
		},
		{
			name: "single element per group",
			input: []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
			},
			keyFn: func(p Person) int {
				return p.Age
			},
			want: map[int][]Person{
				25: {{"Alice", 25}},
				30: {{"Bob", 30}},
				35: {{"Charlie", 35}},
			},
		},
		{
			name:  "empty slice",
			input: []Person{},
			keyFn: func(p Person) int {
				return p.Age
			},
			want: map[int][]Person{},
		},
		{
			name:  "nil slice",
			input: nil,
			keyFn: func(p Person) int {
				return p.Age
			},
			want: map[int][]Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.GroupBy(tt.input, tt.keyFn)

			// Check if maps have same keys
			if len(got) != len(tt.want) {
				t.Errorf("GroupBy() got %d groups, want %d", len(got), len(tt.want))
				return
			}

			// Check each group
			for k, wantGroup := range tt.want {
				gotGroup, exists := got[k]
				if !exists {
					t.Errorf("GroupBy() missing group with key %v", k)
					continue
				}

				if len(gotGroup) != len(wantGroup) {
					t.Errorf("GroupBy() group %v has %d elements, want %d", k, len(gotGroup), len(wantGroup))
					continue
				}

				// Check group contents (may be in different order)
				for _, wantPerson := range wantGroup {
					found := false
					for _, gotPerson := range gotGroup {
						if reflect.DeepEqual(gotPerson, wantPerson) {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("GroupBy() group %v missing person %v", k, wantPerson)
					}
				}
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  0,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := 0
			lang.ForEach(tt.input, func(i int) {
				sum += i
			})

			if sum != tt.want {
				t.Errorf("ForEach() accumulated sum = %v, want %v", sum, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:  "all match",
			input: []int{2, 4, 6, 8},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name:  "some don't match",
			input: []int{2, 4, 5, 8},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
		{
			name:  "none match",
			input: []int{1, 3, 5, 7},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name:  "nil slice",
			input: nil,
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.All(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:  "some match",
			input: []int{1, 4, 5, 7},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name:  "all match",
			input: []int{2, 4, 6, 8},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name:  "none match",
			input: []int{1, 3, 5, 7},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
		{
			name:  "nil slice",
			input: nil,
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Any(tt.input, tt.predicate)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take more than length",
			input: []int{1, 2, 3},
			n:     5,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take zero",
			input: []int{1, 2, 3, 4, 5},
			n:     0,
			want:  []int{},
		},
		{
			name:  "take negative",
			input: []int{1, 2, 3, 4, 5},
			n:     -3,
			want:  []int{},
		},
		{
			name:  "empty slice",
			input: []int{},
			n:     3,
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			n:     3,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Take(tt.input, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}

			// Verify it's a copy, not the original
			if len(tt.input) > 0 && len(got) > 0 {
				// Modify the copy
				got[0] = 999
				if len(tt.input) > 0 && tt.input[0] == 999 {
					t.Errorf("Take() modified the original slice")
				}
			}
		})
	}
}

func TestSkip(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{3, 4, 5},
		},
		{
			name:  "skip exactly length",
			input: []int{1, 2, 3},
			n:     3,
			want:  []int{},
		},
		{
			name:  "skip more than length",
			input: []int{1, 2, 3},
			n:     5,
			want:  []int{},
		},
		{
			name:  "skip zero",
			input: []int{1, 2, 3, 4, 5},
			n:     0,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "skip negative",
			input: []int{1, 2, 3, 4, 5},
			n:     -3,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "empty slice",
			input: []int{},
			n:     3,
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			n:     3,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Skip(tt.input, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Skip() = %v, want %v", got, tt.want)
			}

			// Verify it's a copy, not the original
			if len(tt.input) > 0 && len(got) > 0 {
				// Modify the copy
				got[0] = 999
				if len(tt.input) > 0 && tt.input[0] == 999 {
					t.Errorf("Skip() modified the original slice")
				}
			}
		})
	}
}

func TestCompact(t *testing.T) {
	a, b, c := 1, 2, 3
	tests := []struct {
		name  string
		input []*int
		want  []*int
	}{
		{
			name:  "normal usage",
			input: []*int{&a, nil, &b, nil, &c},
			want:  []*int{&a, &b, &c},
		},
		{
			name:  "no nil values",
			input: []*int{&a, &b, &c},
			want:  []*int{&a, &b, &c},
		},
		{
			name:  "all nil values",
			input: []*int{nil, nil, nil},
			want:  []*int{},
		},
		{
			name:  "empty slice",
			input: []*int{},
			want:  []*int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.Compact(tt.input)

			// Need to compare pointers, not values
			if len(got) != len(tt.want) {
				t.Errorf("Compact() len = %v, want %v", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Compact() at index %d = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestMergeMap(t *testing.T) {
	tests := []struct {
		name string
		maps []map[string]int
		want map[string]int
	}{
		{
			name: "normal usage",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
			},
			want: map[string]int{"a": 1, "b": 3, "c": 4},
		},
		{
			name: "three maps",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
				{"c": 5, "d": 6},
			},
			want: map[string]int{"a": 1, "b": 3, "c": 5, "d": 6},
		},
		{
			name: "with empty map",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{},
				{"c": 3, "d": 4},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "with nil map",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				nil,
				{"c": 3, "d": 4},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "all empty maps",
			maps: []map[string]int{
				{},
				{},
			},
			want: map[string]int{},
		},
		{
			name: "all nil maps",
			maps: []map[string]int{
				nil,
				nil,
			},
			want: map[string]int{},
		},
		{
			name: "no maps",
			maps: []map[string]int{},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.MergeMap(tt.maps...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZipToMap(t *testing.T) {
	tests := []struct {
		name   string
		keys   []string
		values []int
		want   map[string]int
	}{
		{
			name:   "normal usage",
			keys:   []string{"a", "b", "c"},
			values: []int{1, 2, 3},
			want:   map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:   "more keys than values",
			keys:   []string{"a", "b", "c", "d"},
			values: []int{1, 2, 3},
			want:   map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:   "more values than keys",
			keys:   []string{"a", "b"},
			values: []int{1, 2, 3, 4},
			want:   map[string]int{"a": 1, "b": 2},
		},
		{
			name:   "empty keys",
			keys:   []string{},
			values: []int{1, 2, 3},
			want:   map[string]int{},
		},
		{
			name:   "empty values",
			keys:   []string{"a", "b", "c"},
			values: []int{},
			want:   map[string]int{},
		},
		{
			name:   "nil keys",
			keys:   nil,
			values: []int{1, 2, 3},
			want:   map[string]int{},
		},
		{
			name:   "nil values",
			keys:   []string{"a", "b", "c"},
			values: nil,
			want:   map[string]int{},
		},
		{
			name:   "both nil",
			keys:   nil,
			values: nil,
			want:   map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lang.ZipToMap(tt.keys, tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZipToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		wantMatch []int
		wantRest  []int
	}{
		{
			name:  "normal usage",
			input: []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			wantMatch: []int{2, 4, 6},
			wantRest:  []int{1, 3, 5},
		},
		{
			name:  "all match",
			input: []int{2, 4, 6, 8},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			wantMatch: []int{2, 4, 6, 8},
			wantRest:  []int{},
		},
		{
			name:  "none match",
			input: []int{1, 3, 5, 7},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			wantMatch: []int{},
			wantRest:  []int{1, 3, 5, 7},
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			wantMatch: []int{},
			wantRest:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			predicate: func(i int) bool {
				return i%2 == 0
			},
			wantMatch: nil,
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotRest := lang.Partition(tt.input, tt.predicate)

			// Sort for consistent ordering
			sort.Ints(gotMatch)
			sort.Ints(tt.wantMatch)
			sort.Ints(gotRest)
			sort.Ints(tt.wantRest)

			if !reflect.DeepEqual(gotMatch, tt.wantMatch) {
				t.Errorf("Partition() match = %v, want %v", gotMatch, tt.wantMatch)
			}
			if !reflect.DeepEqual(gotRest, tt.wantRest) {
				t.Errorf("Partition() rest = %v, want %v", gotRest, tt.wantRest)
			}
		})
	}
}
