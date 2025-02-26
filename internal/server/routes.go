package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"io/fs"
	"jwt-test/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	staticFiles, _ := fs.Sub(web.Files, "assets")
	r.StaticFS("/assets", http.FS(staticFiles))

	r.GET("/web", func(c *gin.Context) {
		templ.Handler(web.HelloForm()).ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/hello", func(c *gin.Context) {
		web.HelloWebHandler(c.Writer, c.Request)
	})

	r.GET("/login", func(c *gin.Context) {
		templ.Handler(web.LoginPage()).ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/login", HandleLogin)

	r.GET("/", AuthenticateMiddleware, func(c *gin.Context) {
		templ.Handler(web.StartPage()).ServeHTTP(c.Writer, c.Request)
	})

	// Draw
	r.GET("/draw", func(c *gin.Context) {
		web.AllDrawsHandler(c.Writer, c.Request)
	})
	r.GET("/draw/:id", func(c *gin.Context) {
		id := c.Param("id")
		web.DrawHandler(id, c.Writer, c.Request)
	})
	r.GET("/newdraw", func(c *gin.Context) {
		web.NewDrawsHandler(c.Writer, c.Request)
	})
	r.POST("/newdraw", func(c *gin.Context) {
		web.NewDrawFormsHandler(c.Writer, c.Request)
		c.Redirect(http.StatusSeeOther, "/")
	})

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
