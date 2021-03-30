package services

import (
	"github.com/thyago/directed-graph-go/models"
)

type DirectedGraphDAO interface {
	FindAll() ([]models.DirectedGraph, error)
	FindByID(id uint64) (*models.DirectedGraph, error)
	Save(dg *models.DirectedGraph) error
	Delete(id uint64) error
}

type DirectedGraphService struct {
	dao DirectedGraphDAO
}

func NewDirectedGraphService(dao DirectedGraphDAO) *DirectedGraphService {
	return &DirectedGraphService{dao: dao}
}

func (s *DirectedGraphService) Create(name string) (*models.DirectedGraph, error) {
	directedGraph := &models.DirectedGraph{Name: name}
	err := s.dao.Save(directedGraph)
	if err != nil {
		return nil, err
	}
	return directedGraph, nil
}

func (s *DirectedGraphService) List() ([]models.DirectedGraph, error) {
	return s.dao.FindAll()
}

func (s *DirectedGraphService) Update(directedGraphID uint64, name string) error {
	directedGraph, err := s.dao.FindByID(directedGraphID)
	if err != nil {
		return err
	}
	directedGraph.Name = name
	return s.dao.Save(directedGraph)
}

func (s *DirectedGraphService) Get(directedGraphID uint64) (*models.DirectedGraph, error) {
	m, err := s.dao.FindByID(directedGraphID)
	return m, err
}

func (s *DirectedGraphService) Remove(directedGraphID uint64) error {
	return s.dao.Delete(directedGraphID)
}
