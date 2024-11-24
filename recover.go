package lang

import (
	"fmt"
	"runtime/debug"
)

// Logger is an interface for logging panic stack trace and errors.
type Logger interface {
	Error(msg string, args ...any)
}

// Go runs goroutine with recover. It will print stack trace and restart goroutine in case of panic.
// If you want to run goroutine without restarting after panic, just use go func() with Recover.
func Go(l Logger, f func()) {
	var foo func()
	fn := f
	foo = func() {
		defer func() {
			if err := recover(); err != nil {
				printErrorWithStack(l, err)
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
		err := fmt.Errorf("%v", panicErr)
		*outerError = err
		return true
	}
	return false
}

// RecoverWithErrAndStack should be used with defer to recover panic and return error from function with logging stack.
func RecoverWithErrAndStack(l Logger, outerError *error) bool {
	if panicErr := recover(); panicErr != nil {
		err := fmt.Errorf("%v", panicErr)
		*outerError = err
		printErrorWithStack(l, err)
		return true
	}
	return false
}

// RecoverWithHandler should be used with defer to recover and call provided handler func.
func RecoverWithHandler(handler func(err any)) bool {
	if panicErr := recover(); panicErr != nil {
		handler(panicErr)
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
