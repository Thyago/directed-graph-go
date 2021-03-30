//+build test_all unit

package services_test

import (
	"testing"

	"github.com/thyago/directed-graph-go/models"
	"github.com/thyago/directed-graph-go/services"
	"github.com/thyago/directed-graph-go/util"
)

func setupNode() (*MockNodeDAO, *services.NodeService) {
	dgDAO := &MockDirectedGraphDAO{nextID: 1}
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 2, Name: "Test 2"})

	nDAO := &MockNodeDAO{nextID: 1}
	s := services.NewNodeService(nDAO, dgDAO)
	return nDAO, s
}

func Test_CreateNode_WithMetadata(t *testing.T) {
	nDAO, s := setupNode()
	n, err := s.Create(1, map[string]string{"name": "Node 1", "anykey": "anyvalue"})
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := n.ID, nDAO.GetNextID()-1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := len(n.NodeMetadata), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.NodeMetadata[0].Meta, "name"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.NodeMetadata[0].Content, "Node 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.NodeMetadata[1].Meta, "anykey"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.NodeMetadata[1].Content, "anyvalue"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_CreateNode_WithoutMetadata(t *testing.T) {
	nDAO, s := setupNode()
	n, err := s.Create(1, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := n.ID, nDAO.GetNextID()-1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := n.DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := len(n.NodeMetadata), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_ListNode(t *testing.T) {
	nDAO, s := setupNode()

	// Empty
	ns, err := s.List(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(ns), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              1,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 1"}},
	})
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              2,
		DirectedGraphID: 1,
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
	if got, want := nDAO.Nodes[0].ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[0].DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := len(nDAO.Nodes[0].NodeMetadata), 1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[0].Meta, "name"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[0].Content, "Node 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[1].ID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[1].DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := len(nDAO.Nodes[1].NodeMetadata), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_GetNode(t *testing.T) {
	nDAO, s := setupNode()
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              1,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 1"}},
	})

	// Get non existing
	n, err := s.Get(1, 2)
	if err != util.ErrNotFound {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
		return
	}
	if n != nil {
		t.Errorf("Expected nil, got %v", n)
	}

	// Get created
	getN, err := s.Get(1, 1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := getN.ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getN.DirectedGraphID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := len(getN.NodeMetadata), 1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := getN.NodeMetadata[0].Meta, "name"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getN.NodeMetadata[0].Content, "Node 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_UpdateNode(t *testing.T) {
	nDAO, s := setupNode()
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              1,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 1"}},
	})

	// Update
	nodeMetadata := map[string]string{
		"name":   "Node 1 (New Name)",
		"anykey": "Anyvalue",
	}
	err := s.Update(1, 1, nodeMetadata)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	if got, want := len(nDAO.Nodes[0].NodeMetadata), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[0].Meta, "name"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[0].Content, "Node 1 (New Name)"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[1].Meta, "anykey"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := nDAO.Nodes[0].NodeMetadata[1].Content, "Anyvalue"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_RemoveNode(t *testing.T) {
	nDAO, s := setupNode()
	nDAO.Nodes = append(nDAO.Nodes, models.Node{
		ID:              1,
		DirectedGraphID: 1,
		NodeMetadata:    []models.NodeMetadata{models.NodeMetadata{NodeID: 1, Meta: "name", Content: "Node 1"}},
	})

	// Remove
	err := s.Remove(1, 1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	// Get
	if len(nDAO.Nodes) > 0 {
		t.Error("Not removed")
	}
}

// MOCK

type MockNodeDAO struct {
	nextID uint64
	Nodes  []models.Node
}

func (m *MockNodeDAO) FindAll(directedGraphID uint64) ([]models.Node, error) {
	n := []models.Node{}
	for _, element := range m.Nodes {
		if element.DirectedGraphID == directedGraphID {
			n = append(n, element)
		}
	}
	return n, nil
}

func (m *MockNodeDAO) FindByID(directedGraphID, id uint64) (*models.Node, error) {
	for _, element := range m.Nodes {
		if element.DirectedGraphID == directedGraphID && element.ID == id {
			return &element, nil
		}
	}
	return nil, util.ErrNotFound
}

func (m *MockNodeDAO) Save(n *models.Node) error {
	if n.ID != 0 {
		for index, element := range m.Nodes {
			if element.DirectedGraphID == n.DirectedGraphID && element.ID == n.ID {
				m.Nodes[index] = *n
				return nil
			}
		}
		return util.ErrNotFound
	}
	n.ID = m.nextID
	m.nextID++
	m.Nodes = append(m.Nodes, *n)
	return nil
}

func (m *MockNodeDAO) Delete(directedGraphID, id uint64) error {
	for index, element := range m.Nodes {
		if element.DirectedGraphID == directedGraphID && element.ID == id {
			m.Nodes = append(m.Nodes[:index], m.Nodes[index+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockNodeDAO) GetNextID() uint64 {
	return m.nextID
}
