package crow

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func initStore() error {
    return nil
}

func initRouters(g *gin.Engine) error {

    g.POST("/login", func(ctx *gin.Context) {})

    api := g.Group("v1")
    {
        api.GET("/dashborad", func(ctx *gin.Context) {})
    }

    return nil
}

func startSecureHttpService(g *gin.Engine) *http.Server {
    srv := http.Server{
        Addr: ":9527",
        Handler: g,
    }
    return &srv
}
