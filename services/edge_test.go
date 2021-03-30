//+build test_all unit

package services_test

import (
	"testing"

	"github.com/thyago/directed-graph-go/models"
	"github.com/thyago/directed-graph-go/services"
	"github.com/thyago/directed-graph-go/util"
)

func setupEdge() (*MockEdgeDAO, *services.EdgeService) {
	dgDAO := &MockDirectedGraphDAO{nextID: 1}
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 2, Name: "Test 2"})

	nDAO := &MockNodeDAO{nextID: 1}
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              1,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 1"}},
	})
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              2,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 2"}},
	})
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              3,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 3"}},
	})

	eDAO := &MockEdgeDAO{}
	s := services.NewEdgeService(eDAO, dgDAO, nDAO)
	return eDAO, s
}

func Test_CreateEdge_Valid(t *testing.T) {
	_, s := setupEdge()
	e, err := s.Create(1, 1, 2)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := e.DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := e.TailNodeID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := e.HeadNodeID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_CreateEdge_Invalid(t *testing.T) {
	eDAO, s := setupEdge()
	//Create with same node as head and tail
	_, err := s.Create(1, 1, 1)
	if err != util.ErrSelfLoop {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
	}
	if got, want := len(eDAO.Edges), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//Create with invalid head
	_, err = s.Create(1, 5, 1)
	if err != util.ErrNotFound {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
	}
	if got, want := len(eDAO.Edges), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//Create with invalid tail
	_, err = s.Create(1, 2, 4)
	if err != util.ErrNotFound {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
	}
	if got, want := len(eDAO.Edges), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_ListEdge(t *testing.T) {
	eDAO, s := setupEdge()

	// Empty
	ns, err := s.List(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(ns), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// List 2 edges
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      1,
		HeadNodeID:      2,
	})
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      3,
		HeadNodeID:      1,
	})

	ns, err = s.List(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(ns), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := eDAO.Edges[0].TailNodeID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := eDAO.Edges[0].HeadNodeID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := eDAO.Edges[1].TailNodeID, uint64(3); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := eDAO.Edges[1].HeadNodeID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_GetEdge(t *testing.T) {
	eDAO, s := setupEdge()
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      1,
		HeadNodeID:      2,
	})
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      3,
		HeadNodeID:      1,
	})

	// Get non existing
	n, err := s.Get(1, 3, 2)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	} else if err != util.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
		return
	}
	if n != nil {
		t.Errorf("Expected nil, got %v", n)
	}

	// Get created
	getN, err := s.Get(1, 3, 1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := getN.TailNodeID, uint64(3); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getN.HeadNodeID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_RemoveEdge(t *testing.T) {
	eDAO, s := setupEdge()
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      1,
		HeadNodeID:      2,
	})
	eDAO.Edges = append(eDAO.Edges, models.Edge{
		DirectedGraphID: 1,
		TailNodeID:      3,
		HeadNodeID:      1,
	})

	// Remove
	err := s.Remove(1, 3, 1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	// Get
	if len(eDAO.Edges) != 1 {
		t.Error("Not removed")
	}
}

// MOCK

type MockEdgeDAO struct {
	Edges []models.Edge
}

func (m *MockEdgeDAO) FindAll(directedGraphID uint64) ([]models.Edge, error) {
	n := []models.Edge{}
	for _, element := range m.Edges {
		if element.DirectedGraphID == directedGraphID {
			n = append(n, element)
		}
	}
	return n, nil
}

func (m *MockEdgeDAO) FindByPK(directedGraphID, tailNodeID, headNodeID uint64) (*models.Edge, error) {
	for _, element := range m.Edges {
		if element.DirectedGraphID == directedGraphID && element.TailNodeID == tailNodeID && element.HeadNodeID == headNodeID {
			return &element, nil
		}
	}
	return nil, util.ErrNotFound
}

func (m *MockEdgeDAO) Save(n *models.Edge) error {
	m.Edges = append(m.Edges, *n)
	return nil
}

func (m *MockEdgeDAO) Delete(directedGraphID, tailNodeID, headNodeID uint64) error {
	for index, element := range m.Edges {
		if element.DirectedGraphID == directedGraphID && element.TailNodeID == tailNodeID && element.HeadNodeID == headNodeID {
			m.Edges = append(m.Edges[:index], m.Edges[index+1:]...)
			return nil
		}
	}
	return nil
}
