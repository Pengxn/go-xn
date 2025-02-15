package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/util/httplib"
)

var (
	// nightly.link is a service to provide nightly build artifact download link.
	nightlyURL = "https://nightly.link/Pengxn/go-xn/workflows/test/main"

	// Update is "update" subcommand.
	// It's used to update command binary to the latest version.
	updateCmd = &cli.Command{
		Name:   "update",
		Usage:  "Update the binary to the latest version",
		Action: update,
	}
)

func update(c *cli.Context) error {
	link := fmt.Sprintf("%s/%s-%s.zip", nightlyURL, runtime.GOOS, runtime.GOARCH)
	resp, err := httplib.New().GET(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})

	// create real-time progress bar in terminal
	bar := progressbar.DefaultBytes(resp.ContentLength)

	size, err := io.Copy(io.MultiWriter(buff, bar), resp.Body)
	if err != nil {
		return err
	}

	archive, err := zip.NewReader(bytes.NewReader(buff.Bytes()), size)
	if err != nil {
		return err
	}

	return unzip(archive, "build")
}

func unzip(archive *zip.Reader, dst string) error {
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		// create directory if file in archive is a directory
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

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
