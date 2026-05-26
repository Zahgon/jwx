package pool

var errorSlicePool = New[[]error](allocErrorSlice, freeErrorSlice)

func allocErrorSlice() []error { _ = "STUB: not implemented"; return nil }

func freeErrorSlice(s []error) []error {
	_ = "STUB: not implemented"
	// Reset the slice to its zero value
	return nil
}

func ErrorSlice() *Pool[[]error] { _ = "STUB: not implemented"; return nil }
