package handler

import (
	"memrizr/model"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router       *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewServer(us model.UserService, ts model.TokenService) *Server {
	server := &Server{
		UserService:  us,
		TokenService: ts,
	}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	g := router.Group(os.Getenv("ACCOUNT_API_URL"))
	{
		g.GET("/me", server.Me)
		g.POST("/signup", server.Signup)
		g.POST("/signin", server.Signin)
		g.POST("/signout", server.Signout)
		g.POST("/tokens", server.Tokens)
		g.POST("/image", server.Image)
		g.DELETE("/image", server.DeleteImage)
		g.PUT("/details", server.Details)
	}
	server.Router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }
