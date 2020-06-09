package numbers

// MyInt is an alias for int32
type MyInt int32

// Add is bla
func (n MyInt) Add(v int32) MyInt {
	return n + MyInt(v)
}

// Sub is bla
func (n MyInt) Sub(v int32) MyInt {
	return n - MyInt(v)
}

// Multiply is bla
func (n MyInt) Multiply(v int32) MyInt {
	return n * MyInt(v)
}

// Divide is bla
func (n MyInt) Divide(v int32) MyInt {
	return n / MyInt(v)
}
