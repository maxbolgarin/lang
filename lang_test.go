package lang_test

import (
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
