package daos

import (
	"github.com/thyago/directed-graph-go/db"
	"github.com/thyago/directed-graph-go/models"
)

type DirectedGraphDAO struct {
	database *db.Database
}

func NewDirectedGraphDAO(db *db.Database) *DirectedGraphDAO {
	return &DirectedGraphDAO{db}
}

func (dao *DirectedGraphDAO) FindAll() ([]models.DirectedGraph, error) {
	// TODO: Add pagination

	dORM := dao.database.GetORM()
	directedGraphs := []models.DirectedGraph{}
	err := dORM.Debug().Find(&directedGraphs).Limit(100).Error
	if err != nil {
		return nil, err
	}
	return directedGraphs, nil
}

func (dao *DirectedGraphDAO) FindByID(id uint64) (*models.DirectedGraph, error) {
	dORM := dao.database.GetORM()
	directedGraph := &models.DirectedGraph{}
	err := dORM.Debug().Where("id = ?", id).Take(directedGraph).Error
	if err != nil {
		return nil, err
	}
	return directedGraph, nil
}

func (dao *DirectedGraphDAO) Save(dg *models.DirectedGraph) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Save(dg).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *DirectedGraphDAO) Delete(id uint64) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Where("id = ?", id).Delete(&models.DirectedGraph{}).Error
	if err != nil {
		return err
	}
	return nil
}
