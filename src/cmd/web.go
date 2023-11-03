package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/lib/webauthn"
	"github.com/Pengxn/go-xn/src/model"
	"github.com/Pengxn/go-xn/src/route"
	"github.com/Pengxn/go-xn/src/util/log"
)

func init() {
	Web.Flags().IntP("port", "p", 7991, "Temporary port number to prevent conflict")
	Web.Flags().BoolP("debug", "d", false, "Enable debug mode, print extra debugging information")
	Web.Flags().StringP("config", "c", "fyj.ini", "Custom configuration `file` path")

	app.AddCommand(Web)
}

var (
	// Web is "web" subcommand. It's used to run web server.
	Web = &cobra.Command{
		Use:   "web",
		Short: "Start web server interface for blog",
		Long: `Run a performant web server which serves the site for blog.
If '--port' flag is not used, it will use port 7991 by default.`,
		Run: runWeb,
	}
)

// runWeb is the entry point to the web server.
// Parses the arguments, routes and others.
func runWeb(cmd *cobra.Command, args []string) {
	configFile := cmd.Flag("config").Value.String()
	port := cmd.Flag("port").Value.String()
	debug, _ := cmd.Flags().GetBool("debug")

	// Override config by cli flag
	config.OverrideConfigByFlag(configFile, port, debug)
	model.InitTables()
	webauthn.InitWebAuthn()

	err := route.InitRoutes()
	if err != nil {
		log.Fatalln("Fail to Start Web Server.", err)
	}
}
