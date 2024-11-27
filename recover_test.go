package lang_test

import (
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
