package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	app.Version = showVersion()
	app.SetVersionTemplate(`{{ printf "%s" .Version }}`)

	app.AddCommand(VersionCmd)
}

// Version information for cmd
// Use "var" (not const) to defined variable for `go build -ldflags`
// And export 'Version' variable
// ⚠️ WARNING: should only be set by "-ldflags" cli flag.
var (
	Version   = ""
	commitID  = ""
	buildTime = ""

	// VersionCmd is "version" subcommand.
	// It prints the version, revision and built time information to stdout.
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print some information about version",
		Long: `Prints version information that might help you get out of trouble.
And it display revision and built time information.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(showVersion())
		},
	}
)

// showVersion prints the version information to stdout
func showVersion() string {
	return fmt.Sprintf(`FYJ %s
---------------------------------
- Go version: %s
- Revision:   %s
- OS/Arch:    %s/%s
- Built time: %s
`, Version, runtime.Version(), commitID,
		runtime.GOOS, runtime.GOARCH, buildTime)
}
