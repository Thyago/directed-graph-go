package services

import (
	"github.com/thyago/directed-graph-go/models"
)

type NodeDAO interface {
	Save(node *models.Node) error
	Delete(directedGraphID, id uint64) error

	FindAll(directedGraphID uint64) ([]models.Node, error)
	FindByID(directedGraphID, id uint64) (*models.Node, error)
}

type NodeService struct {
	dao   NodeDAO
	dgDAO DirectedGraphDAO
}

func NewNodeService(dao NodeDAO, dgDAO DirectedGraphDAO) *NodeService {
	return &NodeService{dao: dao, dgDAO: dgDAO}
}

func (s *NodeService) Create(directedGraphID uint64, metadata map[string]string) (*models.Node, error) {
	directedGraph, err := s.dgDAO.FindByID(directedGraphID)
	if err != nil {
		return nil, err
	}
	var nodeMetadata []models.NodeMetadata
	for meta, content := range metadata {
		nodeMetadata = append(nodeMetadata, models.NodeMetadata{Meta: meta, Content: content})
	}
	node := models.NewNode(directedGraph, nodeMetadata)
	err = s.dao.Save(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *NodeService) List(directedGraphID uint64) ([]models.Node, error) {
	return s.dao.FindAll(directedGraphID)
}

func (s *NodeService) Update(directedGraphID, nodeID uint64, metadata map[string]string) error {
	node, err := s.dao.FindByID(directedGraphID, nodeID)
	if err != nil {
		return err
	}
	var nodeMetadata []models.NodeMetadata
	for meta, content := range metadata {
		nodeMetadata = append(nodeMetadata, models.NodeMetadata{Meta: meta, Content: content})
	}
	node.NodeMetadata = nodeMetadata
	return s.dao.Save(node)
}

func (s *NodeService) Get(directedGraphID, nodeID uint64) (*models.Node, error) {
	return s.dao.FindByID(directedGraphID, nodeID)
}

func (s *NodeService) Remove(directedGraphID, nodeID uint64) error {
	return s.dao.Delete(directedGraphID, nodeID)
}

func (s *NodeService) ListChildren(directedGraphID, nodeID uint64) ([]models.Node, error) {
	return nil, nil
}

func (s *NodeService) ListParents(directedGraphID, nodeID uint64) ([]models.Node, error) {
	return nil, nil
}
