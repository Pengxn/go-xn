package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-xn/src/route"
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

// Run is the entry point to the server app.
// Parses the arguments routes and others.
func Run() {
	gin.SetMode(gin.ReleaseMode)

	fmt.Printf(banner)

	g := gin.Default()

	route.InitRoutes(g)

	g.Run(":3000")
}
