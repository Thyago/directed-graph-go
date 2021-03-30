package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thyago/directed-graph-go/config"
	"github.com/thyago/directed-graph-go/db"
	"github.com/thyago/directed-graph-go/models"
)

type Server struct {
	router *gin.Engine
	db     *db.Database
	port   string
}

func NewServer() *Server {
	// Load App Config
	config.LoadConfig()

	// Create DB connection
	db := db.NewDatabase(config.Config.DBUser, config.Config.DBPassword, config.Config.DBHost, config.Config.DBPort, config.Config.DBName)

	// Create router
	r := gin.New()

	// Logger middleware
	r.Use(gin.Logger())

	// Recovery middleware: Recover from any panic and returns 500
	r.Use(gin.Recovery())

	return &Server{r, db, config.Config.ServerPort}
}

func (s *Server) Run() {
	err := s.db.Open()
	if err != nil {
		fmt.Printf("Failed to run: %v", err)
	}

	// Migrate DB when needed
	s.db.Migrate(
		&models.DirectedGraph{},
		&models.Node{},
		&models.NodeMetadata{},
		&models.Edge{},
	)

	// Initialize routes
	s.initRouter()
	s.router.Run(fmt.Sprintf(":%v", s.port))
}

func (s *Server) initRouter() {

	// TODO: Include auth

	// Create API v1 routing endpoints
	v1 := s.router.Group("v1")
	{
		v1.POST("/directed-graphs", s.CreateDirectedGraph)
		v1.GET("/directed-graphs", s.ListDirectedGraph)
		v1.PUT("/directed-graphs/:directed-graph-id", s.UpdateDirectedGraph)
		v1.GET("/directed-graphs/:directed-graph-id", s.GetDirectedGraph)
		v1.DELETE("/directed-graphs/:directed-graph-id", s.RemoveDirectedGraph)

		v1.POST("/directed-graphs/:directed-graph-id/nodes", s.CreateNode)
		v1.GET("/directed-graphs/:directed-graph-id/nodes", s.ListNode)
		v1.PUT("/directed-graphs/:directed-graph-id/nodes/:node-id", s.UpdateNode)
		v1.GET("/directed-graphs/:directed-graph-id/nodes/:node-id", s.GetNode)
		v1.DELETE("/directed-graphs/:directed-graph-id/nodes/:node-id", s.RemoveNode)

		v1.POST("/directed-graphs/:directed-graph-id/edges", s.CreateEdge)
		v1.GET("/directed-graphs/:directed-graph-id/edges", s.ListEdge)
		v1.GET("/directed-graphs/:directed-graph-id/edges/:tail-node-id/:head-node-id", s.GetEdge)
		v1.DELETE("/directed-graphs/:directed-graph-id/edges/:tail-node-id/:head-node-id", s.RemoveEdge)

		/*v1.GET("/directed-graphs/:directed-graph-id/nodes/:node-id/children", s.ListNodeChildren)
		v1.GET("/directed-graphs/:directed-graph-id/nodes/:node-id/parents", s.ListNodeParents)*/
	}
}
