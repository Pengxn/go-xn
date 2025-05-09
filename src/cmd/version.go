package cmd

import (
	"context"
	"fmt"
	"runtime"

	"github.com/urfave/cli/v3"
)

// Version information for cmd
// Use "var" (not const) to defined variable for `go build -ldflags`.
// ⚠️ WARNING: should only be set by "-ldflags" cli flag.
var (
	version   = ""
	commitID  = ""
	buildTime = ""

	// versionCmd is "version" subcommand.
	// It prints the version, revision and built time information to stdout.
	versionCmd = &cli.Command{
		Name:  "version",
		Usage: "Print some information about version",
		Description: `Prints version information that might help you get out of trouble.
And it display revision and built time information.`,
		Action: showVersion,
	}
)

// showVersion prints the version information to stdout
func showVersion(ctx context.Context, c *cli.Command) error {
	fmt.Printf(`FYJ %s
---------------------------------
- Go version: %s
- Revision:   %s
- OS/Arch:    %s/%s
- Built time: %s
`, version, runtime.Version(), commitID,
		runtime.GOOS, runtime.GOARCH, buildTime)

	return nil
}
