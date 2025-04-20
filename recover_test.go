package lang_test

import (
	"errors"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/maxbolgarin/lang"
)

type testLogger struct {
	logs    atomic.Int64
	lastMsg atomic.Value
}

func (s *testLogger) Error(msg string, args ...any) {
	s.logs.Add(1)
	s.lastMsg.Store(msg)
}

func TestGo(t *testing.T) {
	var (
		wg         = sync.WaitGroup{}
		l          = testLogger{}
		counter    atomic.Int64
		logCounter = int64(5)
	)

	// Nothing should happen
	lang.Go(nil, nil)

	wg.Add(1)
	lang.Go(&l, func() {
		counter.Add(1)
		if counter.Load() < logCounter {
			panic("panic-error")
		}
		wg.Done()
	})

	wg.Wait()

	if l.logs.Load() != logCounter-1 {
		t.Errorf("expected %d logs", logCounter-1)
	}
}

func TestRecover(t *testing.T) {
	l := testLogger{}
	defer func() {
		if l.logs.Load() == 0 {
			t.Error("expected at least one log")
		}
		if !strings.Contains(l.lastMsg.Load().(string), "lang_test.TestRecover") {
			t.Error("expected stack trace in message")
		}
	}()
	defer lang.Recover(&l)
	panic("panic-error")
}

func TestRecoverNilLog(t *testing.T) {
	defer lang.Recover(nil)
	panic("panic-error")
}

func TestRecoverWithErr(t *testing.T) {
	var err error
	defer func() {
		if err == nil {
			t.Error("expected error")
		}
		if !strings.Contains(err.Error(), "panic-error") {
			t.Error("expected error in message")
		}
	}()
	defer lang.RecoverWithErr(&err)
	panic("panic-error")
}

func TestRecoverWithErrAndStack(t *testing.T) {
	l := testLogger{}
	var err error
	defer func() {
		if l.logs.Load() == 0 {
			t.Error("expected at least one log")
		}
		if !strings.Contains(l.lastMsg.Load().(string), "lang_test.TestRecoverWithErrAndStack") {
			t.Error("expected stack trace in message")
		}
		if err == nil {
			t.Error("expected error")
		}
		if !strings.Contains(err.Error(), "panic-error") {
			t.Error("expected error in message")
		}
	}()
	defer lang.RecoverWithErrAndStack(&l, &err)
	panic("panic-error")
}

func TestRecoverWithHandler(t *testing.T) {
	var counter atomic.Int64
	defer func() {
		if counter.Load() == 0 {
			t.Error("expected at least one log")
		}
	}()
	defer lang.RecoverWithHandler(func(err any) {
		counter.Add(1)
	})
	panic("panic-error")
}

func TestNoPanic(t *testing.T) {
	l := testLogger{}
	var err error
	var counter atomic.Int64
	defer func() {
		if l.logs.Load() != 0 {
			t.Error("not expected logs")
		}
		if err != nil {
			t.Error("not expected error")
		}
		if counter.Load() != 0 {
			t.Error("not expected counter")
		}
	}()

	defer lang.Recover(&l)
	defer lang.RecoverWithErr(&err)
	defer lang.RecoverWithErrAndStack(&l, &err)
	defer lang.RecoverWithHandler(func(err any) {
		counter.Add(1)
	})
}

func TestDefaultIfPanic(t *testing.T) {
	t.Run("string - success", func(t *testing.T) {
		result := lang.DefaultIfPanic("default", func() string {
			return "success"
		})

		if result != "success" {
			t.Errorf("Expected result to be 'success', got %v", result)
		}
	})

	t.Run("string - panic", func(t *testing.T) {
		result := lang.DefaultIfPanic("default", func() string {
			panic("test panic")
		})

		if result != "default" {
			t.Errorf("Expected result to be 'default', got %v", result)
		}
	})

	t.Run("nil function", func(t *testing.T) {
		result := lang.DefaultIfPanic("default", nil)
		if result != "default" {
			t.Errorf("Expected result to be 'default', got %v", result)
		}
	})

	t.Run("int - success", func(t *testing.T) {
		result := lang.DefaultIfPanic(42, func() int {
			return 100
		})

		if result != 100 {
			t.Errorf("Expected result to be 100, got %v", result)
		}
	})

	t.Run("int - panic", func(t *testing.T) {
		result := lang.DefaultIfPanic(42, func() int {
			panic("test panic")
		})

		if result != 42 {
			t.Errorf("Expected result to be 42, got %v", result)
		}
	})

	t.Run("bool - success", func(t *testing.T) {
		result := lang.DefaultIfPanic(false, func() bool {
			return true
		})

		if !result {
			t.Errorf("Expected result to be true, got %v", result)
		}
	})

	t.Run("bool - panic", func(t *testing.T) {
		result := lang.DefaultIfPanic(false, func() bool {
			panic("test panic")
		})

		if result {
			t.Errorf("Expected result to be false, got %v", result)
		}
	})

	t.Run("struct - success", func(t *testing.T) {
		type testStruct struct {
			Value string
			Count int
		}

		defaultValue := testStruct{Value: "default", Count: 0}
		expectedValue := testStruct{Value: "success", Count: 42}

		result := lang.DefaultIfPanic(defaultValue, func() testStruct {
			return expectedValue
		})

		if !reflect.DeepEqual(result, expectedValue) {
			t.Errorf("Expected result to be %+v, got %+v", expectedValue, result)
		}
	})

	t.Run("struct - panic", func(t *testing.T) {
		type testStruct struct {
			Value string
			Count int
		}

		defaultValue := testStruct{Value: "default", Count: 0}

		result := lang.DefaultIfPanic(defaultValue, func() testStruct {
			panic("test panic")
		})

		if !reflect.DeepEqual(result, defaultValue) {
			t.Errorf("Expected result to be %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("slice - success", func(t *testing.T) {
		defaultValue := []string{"default"}
		expectedValue := []string{"success", "values"}

		result := lang.DefaultIfPanic(defaultValue, func() []string {
			return expectedValue
		})

		if !reflect.DeepEqual(result, expectedValue) {
			t.Errorf("Expected result to be %+v, got %+v", expectedValue, result)
		}
	})

	t.Run("slice - panic", func(t *testing.T) {
		defaultValue := []string{"default"}

		result := lang.DefaultIfPanic(defaultValue, func() []string {
			panic("test panic")
		})

		if !reflect.DeepEqual(result, defaultValue) {
			t.Errorf("Expected result to be %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("map - success", func(t *testing.T) {
		defaultValue := map[string]int{"default": 0}
		expectedValue := map[string]int{"success": 42}

		result := lang.DefaultIfPanic(defaultValue, func() map[string]int {
			return expectedValue
		})

		if !reflect.DeepEqual(result, expectedValue) {
			t.Errorf("Expected result to be %+v, got %+v", expectedValue, result)
		}
	})

	t.Run("map - panic", func(t *testing.T) {
		defaultValue := map[string]int{"default": 0}

		result := lang.DefaultIfPanic(defaultValue, func() map[string]int {
			panic("test panic")
		})

		if !reflect.DeepEqual(result, defaultValue) {
			t.Errorf("Expected result to be %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("error - success", func(t *testing.T) {
		defaultValue := errors.New("default error")
		expectedValue := errors.New("success error")

		result := lang.DefaultIfPanic(defaultValue, func() error {
			return expectedValue
		})

		if result.Error() != expectedValue.Error() {
			t.Errorf("Expected result to be %v, got %v", expectedValue, result)
		}
	})

	t.Run("error - panic", func(t *testing.T) {
		defaultValue := errors.New("default error")

		result := lang.DefaultIfPanic(defaultValue, func() error {
			panic("test panic")
		})

		if result.Error() != defaultValue.Error() {
			t.Errorf("Expected result to be %v, got %v", defaultValue, result)
		}
	})

	t.Run("nil - success", func(t *testing.T) {
		var defaultValue *string = nil
		expectedPtr := new(string)
		*expectedPtr = "success"

		result := lang.DefaultIfPanic(defaultValue, func() *string {
			return expectedPtr
		})

		if result != expectedPtr || *result != "success" {
			t.Errorf("Expected result to be %v, got %v", expectedPtr, result)
		}
	})

	t.Run("nil - panic", func(t *testing.T) {
		var defaultValue *string = nil

		result := lang.DefaultIfPanic(defaultValue, func() *string {
			panic("test panic")
		})

		if result != nil {
			t.Errorf("Expected result to be nil, got %v", result)
		}
	})

	t.Run("function with parameters", func(t *testing.T) {
		multiply := func(a, b int) int {
			return a * b
		}

		// We need to wrap the function call in a closure to match the expected signature
		result := lang.DefaultIfPanic(0, func() int {
			return multiply(5, 7)
		})

		if result != 35 {
			t.Errorf("Expected result to be 35, got %v", result)
		}
	})

	t.Run("nested panic recovery", func(t *testing.T) {
		// Test nested DefaultIfPanic calls
		outerResult := lang.DefaultIfPanic("outer default", func() string {
			// This inner call should recover from its panic and return "inner default"
			innerResult := lang.DefaultIfPanic("inner default", func() string {
				panic("inner panic")
			})

			// So this should execute
			if innerResult != "inner default" {
				t.Errorf("Inner result should be 'inner default', got %v", innerResult)
			}

			return "outer success"
		})

		if outerResult != "outer success" {
			t.Errorf("Expected outerResult to be 'outer success', got %v", outerResult)
		}
	})
}
