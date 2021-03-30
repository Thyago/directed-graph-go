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
	nodeService *services.NodeService = nil
)

type NodeResponseData struct {
	ID       uint64            `json:"id"`
	Metadata map[string]string `json:"metadata"`
}

func (s *Server) getNodeService() *services.NodeService {
	if nodeService == nil {
		nodeService = services.NewNodeService(daos.NewNodeDAO(s.db), daos.NewDirectedGraphDAO(s.db))
	}
	return nodeService
}

func (s *Server) CreateNode(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	req := &struct{ Metadata map[string]string }{}
	err = unmarshalBody(c, req)
	if err != nil {
		return
	}

	node, err := s.getNodeService().Create(dgID, req.Metadata)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	rMetadata := make(map[string]string)
	for _, element := range node.NodeMetadata {
		rMetadata[element.Meta] = element.Content
	}
	r := &ResponseData{&NodeResponseData{node.ID, rMetadata}}
	c.JSON(http.StatusCreated, r)
}

func (s *Server) ListNode(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}

	nodes, err := s.getNodeService().List(dgID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	items := make([]interface{}, len(nodes))
	for i := 0; i < len(nodes); i++ {
		rMetadata := make(map[string]string)
		for _, element := range nodes[i].NodeMetadata {
			rMetadata[element.Meta] = element.Content
		}
		items[i] = NodeResponseData{ID: nodes[i].ID, Metadata: rMetadata}
	}
	r := &ResponseDataArray{Data: items}
	c.JSON(http.StatusOK, r)
}

func (s *Server) UpdateNode(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	nID, err := getUrlParamUINT64(c, "node-id")
	if err != nil {
		return
	}
	req := &struct{ Metadata map[string]string }{}
	err = unmarshalBody(c, req)
	if err != nil {
		return
	}

	err = s.getNodeService().Update(dgID, nID, req.Metadata)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetNode(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	nID, err := getUrlParamUINT64(c, "node-id")
	if err != nil {
		return
	}

	node, err := s.getNodeService().Get(dgID, nID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			setResponseError(c, http.StatusNotFound, "Not found", err)
		} else {
			setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		}
		return
	}

	rMetadata := make(map[string]string)
	for _, element := range node.NodeMetadata {
		rMetadata[element.Meta] = element.Content
	}
	r := &ResponseData{&NodeResponseData{node.ID, rMetadata}}
	c.JSON(http.StatusOK, r)
}

func (s *Server) RemoveNode(c *gin.Context) {
	dgID, err := getUrlParamUINT64(c, "directed-graph-id")
	if err != nil {
		return
	}
	nID, err := getUrlParamUINT64(c, "node-id")
	if err != nil {
		return
	}

	err = s.getNodeService().Remove(dgID, nID)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

/*func (s *Server) ListNodeChildren(c *gin.Context) {
	// TODO
}

func (s *Server) ListNodeParents(c *gin.Context) {
	// TODO
}*/
