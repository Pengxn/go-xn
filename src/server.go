package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/route"
)

// Information for Go-xn
const (
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
// Parses the arguments, routes and others.
func Run() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	route.InitRoutes(g)

	if err := g.Run(":3000"); err != nil {
		log.Fatalln("Fail to Start Web Server.", err.Error())
		// Exit web server
		os.Exit(1)
	}
}
