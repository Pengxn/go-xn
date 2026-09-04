package compress

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Unzip extracts a .zip archive from the provided reader to the specified destination directory.
func Unzip(r io.Reader, dst string) error {
	// Read all data from reader
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		// create directory if file in archive is a directory
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		// create parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}
		defer fileInArchive.Close()

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
	}

	return nil
}
