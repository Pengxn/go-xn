package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

// Version information for cmd
// Use "var" (not const) to defined variable for `go build -ldflags`
// And export 'Version' variable
// ⚠️ WARNING: should only be set by "-ldflags" cli flag.
var (
	Version   = ""
	commitID  = ""
	buildDate = ""
	buildTime = ""

	// VersionCmd is "version" subcommand.
	// It prints the version, revision and buildt time informations to stdout.
	VersionCmd = &cli.Command{
		Name:  "version",
		Usage: "Print some information about version",
		Description: `Prints version information that might help you get out of trouble.
And it display revision and built time informations.`,
		Action: showVersion,
	}
)

// showVersion prints the version information to stdout
func showVersion(c *cli.Context) error {
	fmt.Printf("FYJ  %s\n", Version)
	fmt.Println("---------------------------------")
	fmt.Printf("- OS/Arch:    %s %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- Go version: %s\n", runtime.Version())
	fmt.Printf("- Revision:   %s\n", commitID)
	fmt.Printf("- Built time: %s %s\n", buildDate, buildTime)

	return nil
}
