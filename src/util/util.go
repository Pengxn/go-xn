package util

// PanicIf throw an error, if it exists
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
