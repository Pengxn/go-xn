package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-xn/src"
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
	fmt.Printf(banner)

	gin.SetMode(gin.ReleaseMode)

	app.Run()
}
