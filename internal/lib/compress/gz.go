package compress

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// Ungzip extracts a .tar.gz archive from the provided reader to the specified destination directory.
func Ungzip(r io.Reader, dst string) error {
	// Read all data from reader
	data, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer data.Close()

	// create tar reader
	archives := tar.NewReader(data)

	for {
		header, err := archives.Next()
		if err == io.EOF { // no more files
			break
		}
		if err != nil {
			return err
		}

		filePath := filepath.Join(dst, header.Name)

		// create directory if entry is a directory
		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// create parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		// create and write file
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, archives); err != nil {
			return err
		}
	}

	return nil
}
