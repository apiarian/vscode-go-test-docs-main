// Package vendoroverride in the gopath defines foo and bar with some docs.
// These have been overridden in the vendor directory of the main package. We
// could probably do the same with branches or something, but this seems a bit
// more controllable.
package vendoroverride

// Foo returns the string "bar" (the vendor version)
func Foo() string {
	return "bar"
}

// Bar returns the string "foo" (the vendor version)
func Bar() string {
	return "foo"
}
