package cmd

import (
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/util/httplib"
)

var (
	nightlyURL = "https://nightly.link/Pengxn/go-xn/workflows/test/main/linux-amd64.zip"

	// Update is "update" subcommand.
	// It's used to update command binary to the latest version.
	Update = &cli.Command{
		Name:   "update",
		Usage:  "Update the binary to the latest version",
		Action: update,
		Flags: []cli.Flag{
			configFile,
		},
	}
)

func update(c *cli.Context) error {
	resp, err := httplib.New().GET(nightlyURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create and write to file
	f, err := os.OpenFile("linux-amd64.zip", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// create real-time progress bar in terminal
	bar := progressbar.DefaultBytes(resp.ContentLength)

	_, err = io.Copy(io.MultiWriter(bar), resp.Body)

	return err
}
