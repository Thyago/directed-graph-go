package services

import (
	"github.com/thyago/directed-graph-go/models"
	"github.com/thyago/directed-graph-go/util"
)

type EdgeDAO interface {
	Save(node *models.Edge) error
	Delete(directedGraphID, tailNodeID, headNodeID uint64) error

	FindAll(directedGraphID uint64) ([]models.Edge, error)
	FindByPK(directedGraphID, tailNodeID, headNodeID uint64) (*models.Edge, error)
}

type EdgeService struct {
	dao   EdgeDAO
	dgDAO DirectedGraphDAO
	nDAO  NodeDAO
}

func NewEdgeService(dao EdgeDAO, dgDAO DirectedGraphDAO, nDAO NodeDAO) *EdgeService {
	return &EdgeService{dao: dao, dgDAO: dgDAO, nDAO: nDAO}
}

func (s *EdgeService) Create(directedGraphID, tailNodeID, headNodeID uint64) (*models.Edge, error) {
	directedGraph, err := s.dgDAO.FindByID(directedGraphID)
	if err != nil {
		return nil, err
	}
	if tailNodeID == headNodeID {
		return nil, util.ErrSelfLoop
	}

	tailNode, err := s.nDAO.FindByID(directedGraphID, tailNodeID)
	if err != nil {
		return nil, err
	}
	headNode, err := s.nDAO.FindByID(directedGraphID, headNodeID)
	if err != nil {
		return nil, err
	}

	//Check if edge already exists
	edge, err := s.dao.FindByPK(directedGraph.ID, tailNode.ID, headNode.ID)
	if err == nil {
		return nil, util.ErrAlreadyExists
	} else if err != util.ErrNotFound {
		return nil, err
	}

	edge = models.NewEdge(directedGraph, tailNode, headNode)
	err = s.dao.Save(edge)
	if err != nil {
		return nil, err
	}
	return edge, nil
}

func (s *EdgeService) List(directedGraphID uint64) ([]models.Edge, error) {
	return s.dao.FindAll(directedGraphID)
}

func (s *EdgeService) Get(directedGraphID, tailNodeID, headNodeID uint64) (*models.Edge, error) {
	return s.dao.FindByPK(directedGraphID, tailNodeID, headNodeID)
}

func (s *EdgeService) Remove(directedGraphID, tailNodeID, headNodeID uint64) error {
	return s.dao.Delete(directedGraphID, tailNodeID, headNodeID)
}
