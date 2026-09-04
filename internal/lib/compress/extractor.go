package compress

import (
	"errors"
	"io"
)

// ExtractFunc is a function type that defines the signature for extraction functions.
// For different archive formats, we can implement this interface to provide a unified extraction mechanism.
// The following functions implement the ExtractFunc interface for .zip and .tar.gz formats, respectively.
type ExtractFunc func(r io.Reader, dst string) error

// Unsupported is a placeholder function for unsupported archive formats.
// It returns an error indicating that the archive format is not supported.
func Unsupported(io.Reader, string) error {
	return errors.New("unsupported archive format")
}
