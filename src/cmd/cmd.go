package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Pengxn/go-xn/src/util/log"
)

var app = &cobra.Command{
	Use:   "go-xn",
	Short: "The platform for publishing and running your blog",
}

// Execute to run cmd
func Execute() {
	if err := app.Execute(); err != nil {
		log.Fatalln("Fail to start app...", err)
	}
}
