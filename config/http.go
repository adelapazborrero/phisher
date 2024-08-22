package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHTTP() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func RegisterRoutes(engine *gin.Engine, routes ...[]Route) {
	for _, routeSlice := range routes {
		for _, route := range routeSlice {
			switch route.Method {
			case "GET":
				engine.GET(route.Path, route.Handler)
			case "POST":
				engine.POST(route.Path, route.Handler)
			case "PUT":
				engine.PUT(route.Path, route.Handler)
			case "DELETE":
				engine.DELETE(route.Path, route.Handler)
			}
		}
	}
}

type Server struct {
	Port   string
	Host   string
	Engine *gin.Engine
}

func NewServer(port, host string) *Server {
	return &Server{
		Port:   port,
		Host:   host,
		Engine: gin.Default(),
	}
}

func (s *Server) Start() {
	s.Engine.Run(s.Host + s.Port)
}
