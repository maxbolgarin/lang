package lang

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// Logger is an interface for logging panic stack trace and errors.
type Logger interface {
	Error(msg string, args ...any)
}

// Go runs goroutine with recover. It will print stack trace and restart goroutine in case of panic.
// If you want to run goroutine without restarting after panic, just use go func() with Recover.
func Go(l Logger, f func()) {
	if f == nil {
		return
	}

	var restartLock sync.Mutex
	var lastRestart time.Time
	maxRestartsPerMinute := 60 // Limit to 1 restart per second on average

	var foo func()
	fn := f
	foo = func() {
		defer func() {
			if err := recover(); err != nil {
				printErrorWithStack(l, err)

				// Calculate restart delay based on previous panics
				restartLock.Lock()
				now := time.Now()
				elapsed := now.Sub(lastRestart)
				var delay time.Duration

				// If we've had a restart recently, add a delay to avoid thrashing
				if elapsed < time.Minute && !lastRestart.IsZero() {
					minInterval := time.Minute / time.Duration(maxRestartsPerMinute)
					if elapsed < minInterval {
						delay = minInterval - elapsed
					}
				}

				// Update the last restart time and unlock
				lastRestart = now
				restartLock.Unlock()

				// Apply delay if needed
				if delay > 0 {
					time.Sleep(delay)
				}

				// Restart the goroutine
				go foo()
			}
		}()
		fn()
	}
	go foo()
}

// Recover should be used with defer to recover and log stack trace in case of panic.
func Recover(l Logger) bool {
	if err := recover(); err != nil {
		printErrorWithStack(l, err)
		return true
	}
	return false
}

// RecoverWithErr should be used with defer to recover panic and return error from function without logging.
func RecoverWithErr(outerError *error) bool {
	if panicErr := recover(); panicErr != nil {
		if outerError != nil {
			*outerError = fmt.Errorf("%v", panicErr)
		}
		return true
	}
	return false
}

// RecoverWithErrAndStack should be used with defer to recover panic and return error from function with logging stack.
func RecoverWithErrAndStack(l Logger, outerError *error) bool {
	if panicErr := recover(); panicErr != nil {
		err := fmt.Errorf("%v", panicErr)
		if outerError != nil {
			*outerError = err
		}
		printErrorWithStack(l, err)
		return true
	}
	return false
}

// RecoverWithHandler should be used with defer to recover and call provided handler func.
func RecoverWithHandler(handler func(err any)) bool {
	if panicErr := recover(); panicErr != nil {
		if handler != nil {
			handler(panicErr)
		}
		return true
	}
	return false
}

func printErrorWithStack(l Logger, err any) {
	if l == nil {
		return
	}
	stack := debug.Stack()
	l.Error(string(stack), "error", err) // build with -trimpath to avoid printing build path in trace
}

// DefaultIfPanic returns the result of the function, or the default value if the function panics.
//
//	result := DefaultIfPanic("default", func() string {
//	    // operation that might panic
//	    return "success"
//	})
func DefaultIfPanic[T any](defaultValue T, f func() T) (result T) {
	if f == nil {
		return defaultValue
	}

	defer func() {
		if r := recover(); r != nil {
			result = defaultValue
		}
	}()
	return f()
}
