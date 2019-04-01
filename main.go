package main

import (
	"log"

	"go-xn/src"

	"github.com/gin-gonic/gin"
)

// Information for Go-xn
const (
	Version = "1.0.0"
	URL     = "https://xn--02f.com"
	banner  = `
    ____
   / ___| ___        __  ___ __
  | |  _ / _ \  ____ \ \/ / '_ \
  | |_| | (_) ||____| >  <| | | |
   \____|\___/       /_/\_\_| |_|
`
)

func main() {
	log.Printf(banner)

	gin.SetMode(gin.ReleaseMode)

	app.Run()
}
