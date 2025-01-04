package route

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// debugRoutes registers debug routes, such as pprof.
func debugRoutes(g gin.IRouter) {
	// Register pprof handlers, refer to:
	// https://golang.org/pkg/net/http/pprof/
	p := g.Group("/debug/pprof")
	p.GET("/", gin.WrapF(pprof.Index))
	p.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	p.GET("/profile", gin.WrapF(pprof.Profile))
	p.POST("/symbol", gin.WrapF(pprof.Symbol))
	p.GET("/symbol", gin.WrapF(pprof.Symbol))
	p.GET("/trace", gin.WrapF(pprof.Trace))
	p.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
	p.GET("/block", gin.WrapH(pprof.Handler("block")))
	p.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	p.GET("/heap", gin.WrapH(pprof.Handler("heap")))
	p.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
	p.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
}
