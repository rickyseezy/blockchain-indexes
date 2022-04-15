package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/rickyseezy/block/docs"
	"github.com/rickyseezy/block/internal/interfaces/rest/middlewares"
	"github.com/rickyseezy/block/internal/usecases"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type ServerEngine interface {
	Start() error
}

type Server struct {
	router        *gin.Engine
	port          string
	blockIndexApp usecases.BlockIndex
}

func NewServer(r *gin.Engine, bApp usecases.BlockIndex, p string) *Server {
	return &Server{
		router:        r,
		port:          p,
		blockIndexApp: bApp,
	}
}

func (s *Server) configure() {
	groups := s.router.Group("/groups")
	{
		groups.GET("/", s.groups)
		groups.GET("/:id", s.group)
	}
	s.router.GET("/indexes/:id", s.index)
	s.router.GET("/blocks/:search", s.block)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.router.Use(middlewares.Cors())
}

func (s *Server) Start() error {
	s.configure()

	err := s.router.Run(fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}

	return nil
}
