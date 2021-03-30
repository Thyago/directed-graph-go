package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thyago/directed-graph-go/daos"
	"github.com/thyago/directed-graph-go/services"
	"github.com/thyago/directed-graph-go/util"
)

var (
	directedGraphService *services.DirectedGraphService = nil
)

type DirectedGraphResponseData struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (s *Server) getDirectedGraphService() *services.DirectedGraphService {
	if directedGraphService == nil {
		directedGraphService = services.NewDirectedGraphService(daos.NewDirectedGraphDAO(s.db))
	}
	return directedGraphService
}

func (s *Server) CreateDirectedGraph(c *gin.Context) {
	req := &struct{ Name string }{}
	err := unmarshalBody(c, req)
	if err != nil {
		return
	}

	directedGraph, err := s.getDirectedGraphService().Create(req.Name)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	r := ResponseData{DirectedGraphResponseData{directedGraph.ID, directedGraph.Name}}
	c.JSON(http.StatusCreated, r)
}

func (s *Server) ListDirectedGraph(c *gin.Context) {
	directedGraphs, err := s.getDirectedGraphService().List()
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	items := make([]interface{}, len(directedGraphs))
	for i := 0; i < len(directedGraphs); i++ {
		items[i] = DirectedGraphResponseData{directedGraphs[i].ID, directedGraphs[i].Name}
	}
	r := &ResponseDataArray{Data: items}
	c.JSON(http.StatusOK, r)
}

func (s *Server) UpdateDirectedGraph(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	req := &struct{ Name string }{}
	err = unmarshalBody(c, req)
	if err != nil {
		return
	}

	err = s.getDirectedGraphService().Update(dgID, req.Name)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Failed", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetDirectedGraph(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	directedGraph, err := s.getDirectedGraphService().Get(dgID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			setResponseError(c, http.StatusNotFound, "Not found", err)
		} else {
			setResponseError(c, http.StatusBadRequest, "Failed", err)
		}
		return
	}

	r := &ResponseData{&DirectedGraphResponseData{directedGraph.ID, directedGraph.Name}}
	c.JSON(http.StatusOK, r)
}

func (s *Server) RemoveDirectedGraph(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	err = s.getDirectedGraphService().Remove(dgID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Failed", err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
