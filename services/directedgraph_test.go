//+build test_all unit

package services_test

import (
	"testing"

	"github.com/thyago/directed-graph-go/models"
	"github.com/thyago/directed-graph-go/services"
	"github.com/thyago/directed-graph-go/util"
)

func setupDirectedGraph() (*MockDirectedGraphDAO, *services.DirectedGraphService) {
	dgDAO := &MockDirectedGraphDAO{nextID: 1}
	s := services.NewDirectedGraphService(dgDAO)
	return dgDAO, s
}

func Test_CreateDirectedGraph(t *testing.T) {
	dgDAO, s := setupDirectedGraph()
	dg, err := s.Create("Service 1")
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := dg.ID, dgDAO.GetNextID()-1; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := dg.Name, "Service 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_ListDirectedGraph(t *testing.T) {
	dgDAO, s := setupDirectedGraph()

	// Empty
	dgs, err := s.List()
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(dgs), 0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 2, Name: "Test 2"})
	dgs, err = s.List()
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := len(dgs), 2; got != want {
		t.Errorf("got %v, wanted %v", got, want)
		return
	}
	if got, want := dgDAO.directedGraphs[0].ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := dgDAO.directedGraphs[0].Name, "Test 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := dgDAO.directedGraphs[1].ID, uint64(2); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := dgDAO.directedGraphs[1].Name, "Test 2"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_GetDirectedGraph(t *testing.T) {
	dgDAO, s := setupDirectedGraph()
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})

	// Get non existing
	dg, err := s.Get(2)
	if err != util.ErrNotFound {
		t.Errorf("Expected %v, got %v", util.ErrSelfLoop, err)
		return
	}
	if dg != nil {
		t.Errorf("Expected nil, got %v", dg)
	}

	// Get created
	getDG, err := s.Get(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}
	if got, want := getDG.ID, uint64(1); got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := getDG.Name, "Test 1"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_UpdateDirectedGraph(t *testing.T) {
	dgDAO, s := setupDirectedGraph()
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})

	// Update
	err := s.Update(1, "Test 2")
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	if got, want := dgDAO.directedGraphs[0].Name, "Test 2"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_RemoveDirectedGraph(t *testing.T) {
	dgDAO, s := setupDirectedGraph()
	dgDAO.directedGraphs = append(dgDAO.directedGraphs, models.DirectedGraph{ID: 1, Name: "Test 1"})

	// Remove
	err := s.Remove(1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		return
	}

	// Get
	if len(dgDAO.directedGraphs) > 0 {
		t.Error("Not removed")
	}
}

// MOCK

type MockDirectedGraphDAO struct {
	nextID         uint64
	directedGraphs []models.DirectedGraph
}

func (m *MockDirectedGraphDAO) FindAll() ([]models.DirectedGraph, error) {
	return m.directedGraphs, nil
}

func (m *MockDirectedGraphDAO) FindByID(id uint64) (*models.DirectedGraph, error) {
	for _, element := range m.directedGraphs {
		if element.ID == id {
			return &element, nil
		}
	}
	return nil, util.ErrNotFound
}

func (m *MockDirectedGraphDAO) Save(dg *models.DirectedGraph) error {
	if dg.ID != 0 {
		for index, element := range m.directedGraphs {
			if element.ID == dg.ID {
				m.directedGraphs[index] = *dg
				return nil
			}
		}
		return util.ErrNotFound
	}
	dg.ID = m.nextID
	m.nextID++
	m.directedGraphs = append(m.directedGraphs, *dg)
	return nil
}

func (m *MockDirectedGraphDAO) Delete(id uint64) error {
	for index, element := range m.directedGraphs {
		if element.ID == id {
			m.directedGraphs = append(m.directedGraphs[:index], m.directedGraphs[index+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockDirectedGraphDAO) GetNextID() uint64 {
	return m.nextID
}
