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
	edgeService *services.EdgeService = nil
)

type EdgeResponseData struct {
	TailNodeID uint64 `json:"tail_node_id"`
	HeadNodeID uint64 `json:"head_node_id"`
}

func (s *Server) getEdgeService() *services.EdgeService {
	if edgeService == nil {
		edgeService = services.NewEdgeService(daos.NewEdgeDAO(s.db), daos.NewDirectedGraphDAO(s.db), daos.NewNodeDAO(s.db))
	}
	return edgeService
}

func (s *Server) CreateEdge(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	type EdgeData EdgeResponseData
	req := &EdgeData{}
	err = unmarshalBody(c, req)
	if err != nil {
		return
	}

	edge, err := s.getEdgeService().Create(dgID, req.TailNodeID, req.HeadNodeID)
	if err != nil {
		if err == util.ErrAlreadyExists {
			setResponseError(c, http.StatusConflict, "Edge already exists", err)
		} else if err == util.ErrSelfLoop {
			setResponseError(c, http.StatusConflict, "Tail and Head nodes must be different", err)
		} else {
			setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		}
		return
	}
	r := &ResponseData{&EdgeResponseData{edge.TailNodeID, edge.HeadNodeID}}
	c.JSON(http.StatusCreated, r)
}

func (s *Server) ListEdge(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	edges, err := s.getEdgeService().List(dgID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	items := make([]interface{}, len(edges))
	for i := 0; i < len(edges); i++ {
		items[i] = EdgeResponseData{edges[i].TailNodeID, edges[i].HeadNodeID}
	}

	r := &ResponseDataArray{Data: items}
	c.JSON(http.StatusOK, r)
}

func (s *Server) GetEdge(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	tailNodeID, err := getUrlParamUINT64(c, "tail-node-id")
	if err != nil {
		return
	}
	headNodeID, err := getUrlParamUINT64(c, "head-node-id")
	if err != nil {
		return
	}

	edge, err := s.getEdgeService().Get(dgID, tailNodeID, headNodeID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			setResponseError(c, http.StatusNotFound, "Not found", err)
		} else {
			setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		}
		return
	}

	r := &ResponseData{EdgeResponseData{edge.TailNodeID, edge.HeadNodeID}}
	c.JSON(http.StatusOK, r)
}

func (s *Server) RemoveEdge(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	tailNodeID, err := getUrlParamUINT64(c, "tail-node-id")
	if err != nil {
		return
	}
	headNodeID, err := getUrlParamUINT64(c, "head-node-id")
	if err != nil {
		return
	}

	err = s.getEdgeService().Remove(dgID, tailNodeID, headNodeID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
