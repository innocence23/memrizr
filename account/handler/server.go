package handler

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	g := router.Group(os.Getenv("ACCOUNT_API_URL"))
	{
		h := Account{}
		g.GET("/me", h.Me)
		g.POST("/signup", h.Signup)
		g.POST("/signin", h.Signin)
		g.POST("/signout", h.Signout)
		g.POST("/tokens", h.Tokens)
		g.POST("/image", h.Image)
		g.DELETE("/image", h.DeleteImage)
		g.PUT("/details", h.Details)
	}
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }
