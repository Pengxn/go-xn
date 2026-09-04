package compress

import "io"

// ExtractFunc is a function type that defines the signature for extraction functions.
// For different archive formats, we can implement this interface to provide a unified extraction mechanism.
// The following functions implement the ExtractFunc interface for .zip and .tar.gz formats, respectively.
type ExtractFunc func(r io.Reader, dst string) error
