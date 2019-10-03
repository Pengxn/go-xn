package app

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/route"
)

// Information for Go-xn
const (
	Version = "0.0.0-Beta"
	URL     = "https://xn--02f.com"
	Banner  = "\n" +
		"  ,----..\n" +
		" /   /   \\\n" +
		"|   :     :    ,---.      ,---,.                 ,---,\n" +
		".   |  ;. /   '   ,'\\   ,'  .' |,--,  ,--,   ,-+-. /  |\n" +
		".   ; /--`   /   /   |,---.'   ,|'. \\/ .`|  ,--.'|'   |\n" +
		";   | ;  __ .   ; ,. :|   |    |'  \\/  / ; |   |  ,\"' |\n" +
		"|   : |.' .''   | |: ::   :  .'  \\  \\.' /  |   | /  | |\n" +
		".   | '_.' :'   | .; ::   |.'     \\  ;  ;  |   | |  | |\n" +
		"'   ; : \\  ||   :    |`---'      / \\  \\  \\ |   | |  |/\n" +
		"'   | '/  .' \\   \\  /          ./__;   ;  \\|   | |--'\n" +
		"|   :    /    `----'           |   :/\\  \\ ;|   |/\n" +
		" \\   \\ .'                      `---'  `--` '---'\n" +
		"  `---`"
)

// Run is the entry point to the server app.
// Parses the arguments routes and others.
func Run() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	route.InitRoutes(g)

	g.Run(":3000")
}
