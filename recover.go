package lang

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// Logger is an interface for logging panic stack traces and errors.
// It provides a method to log error messages with optional arguments.
type Logger interface {
	Error(msg string, args ...any)
}

// Go runs a goroutine with automatic panic recovery and restart capability.
// If the goroutine panics, it logs the stack trace and restarts the goroutine.
// It includes rate limiting to prevent excessive restarts (max 60 per minute).
// Use this when you want automatic recovery and restart after panics.
//
//	// Example usage:
//	Go(logger, func() {
//	    // Your goroutine code here
//	    for {
//	        // Do work that might panic
//	        time.Sleep(time.Second)
//	    }
//	})
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

// Recover should be used with defer to recover from panics and log the stack trace.
// It returns true if a panic was recovered, false otherwise.
// Use this when you want to handle panics gracefully without stopping execution.
//
//	func riskyOperation() {
//	    defer func() {
//	        if Recover(logger) {
//	            // Handle the panic case
//	        }
//	    }()
//	    // Code that might panic
//	}
func Recover(l Logger) bool {
	if err := recover(); err != nil {
		printErrorWithStack(l, err)
		return true
	}
	return false
}

// RecoverWithErr should be used with defer to recover from panics and convert them to errors.
// It returns true if a panic was recovered, false otherwise.
// The panic is converted to an error and stored in the provided error pointer.
// Use this when you want to convert panics to errors without logging.
//
//	func riskyOperation() (err error) {
//	    defer func() {
//	        if RecoverWithErr(&err) {
//	            // Panic was converted to error
//	        }
//	    }()
//	    // Code that might panic
//	    return nil
//	}
func RecoverWithErr(outerError *error) bool {
	if panicErr := recover(); panicErr != nil {
		if outerError != nil {
			*outerError = fmt.Errorf("%v", panicErr)
		}
		return true
	}
	return false
}

// RecoverWithErrAndStack should be used with defer to recover from panics, convert them to errors, and log the stack trace.
// It returns true if a panic was recovered, false otherwise.
// The panic is converted to an error and stored in the provided error pointer, and the stack trace is logged.
// Use this when you want to convert panics to errors and also log the stack trace.
//
//	func riskyOperation() (err error) {
//	    defer func() {
//	        if RecoverWithErrAndStack(logger, &err) {
//	            // Panic was converted to error and logged
//	        }
//	    }()
//	    // Code that might panic
//	    return nil
//	}
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

// RecoverWithHandler should be used with defer to recover from panics and call a custom handler function.
// It returns true if a panic was recovered, false otherwise.
// The handler function is called with the panic value if a panic occurs.
// Use this when you want custom handling of panics.
//
//	func riskyOperation() {
//	    defer func() {
//	        if RecoverWithHandler(func(panicValue any) {
//	            fmt.Printf("Panic recovered: %v\n", panicValue)
//	        }) {
//	            // Panic was handled
//	        }
//	    }()
//	    // Code that might panic
//	}
func RecoverWithHandler(handler func(err any)) bool {
	if panicErr := recover(); panicErr != nil {
		if handler != nil {
			handler(panicErr)
		}
		return true
	}
	return false
}

// printErrorWithStack logs an error with its stack trace using the provided logger.
// This is a helper function used internally by other recovery functions.
func printErrorWithStack(l Logger, err any) {
	if l == nil {
		return
	}
	stack := debug.Stack()
	l.Error(string(stack), "error", err) // build with -trimpath to avoid printing build path in trace
}

// DefaultIfPanic executes a function and returns its result, or returns a default value if the function panics.
// This is useful for operations that might panic but you want to provide a fallback value.
//
//	result := DefaultIfPanic("default", func() string {
//	    return riskyOperation() // might panic
//	}) // result == "default" if riskyOperation panics, otherwise the actual result
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
