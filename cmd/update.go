package cmd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v3"

	"github.com/Pengxn/go-xn/internal/lib/compress"
	"github.com/Pengxn/go-xn/internal/lib/github"
	"github.com/Pengxn/go-xn/internal/util/httplib"
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
		link, err = github.GetNightlyLink(ctx)
	} else {
		link, err = github.GetLatestAssetLink(ctx)
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

	var extract compress.ExtractFunc

	// detect content type, based on the head of the file buffer
	// refer to file magic numbers, https://en.wikipedia.org/wiki/List_of_file_signatures
	contentType := http.DetectContentType(buff.Bytes())
	switch contentType {
	case "application/x-gzip":
		extract = compress.Ungzip
	case "application/zip":
		extract = compress.Unzip
	default:
		extract = compress.Unsupported
	}

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	return extract(buff, filepath.Dir(exePath))
}
