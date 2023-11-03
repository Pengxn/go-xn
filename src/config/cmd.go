package config

import (
	"os"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/log"
)

func OverrideConfigByFlag(config, port string, debug bool) {
	if config != "" { // specified config file
		f := config
		if !file.IsExist(f) || !file.IsFile(f) {
			log.Error("Specified config file not found: " + f)
		}
		if err := loadConfig(f); err != nil {
			log.Errorf("Load specified config file failed, %+v", err)
		}
	}
	// Server config
	if port != "" {
		Config.Server.Port = port
	}
	Config.Server.Debug = debug
}

func getConfigPathByFlag() string {
	if args := os.Args; len(args) > 0 {
		for k, v := range args {
			if v == "-c" || v == "--config" {
				// Load config file from command line
				if len(args) > k+1 {
					return args[k+1]
				}
			}
		}
	}

	return ""
}
