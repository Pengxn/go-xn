package cmd

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v3"

	"github.com/Pengxn/go-xn/src/lib/github"
	"github.com/Pengxn/go-xn/src/util/httplib"
)

var (
	// Update is "update" subcommand.
	// It's used to update command binary to the latest version.
	updateCmd = &cli.Command{
		Name:   "update",
		Usage:  "Update the binary to the latest version",
		Action: update,
		Flags:  []cli.Flag{nightlyFlag},
	}

	// nightlyFlag is a flag to specify updating to the latest nightly build.
	nightlyFlag = &cli.BoolFlag{
		Name:    "nightly",
		Aliases: []string{"n"},
		Usage:   "Update to the latest nightly build",
	}
)

func update(ctx context.Context, c *cli.Command) error {
	var (
		link string
		err  error
	)
	if c.Bool("nightly") {
		link, err = github.GetNightlyLink()
	} else {
		link, err = github.GetLatestAssetLink()
	}
	if err != nil {
		return err
	}

	resp, err := httplib.New().GET(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})

	// create real-time progress bar in terminal
	bar := progressbar.DefaultBytes(resp.ContentLength)

	_, err = io.Copy(io.MultiWriter(buff, bar), resp.Body)
	if err != nil {
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	return unzip(buff, filepath.Dir(exePath))
}

func unzip(r io.Reader, dst string) error {
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

func ungzip(r io.Reader, dst string) error {
	// Read all data from reader
	data, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer data.Close()

	// create tar reader
	archive := tar.NewReader(data)

	for {
		header, err := archive.Next()
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

		if _, err := io.Copy(dstFile, archive); err != nil {
			return err
		}
	}

	return nil
}
