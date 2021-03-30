package daos

import (
	"github.com/thyago/directed-graph-go/db"
	"github.com/thyago/directed-graph-go/models"
)

type EdgeDAO struct {
	database *db.Database
}

func NewEdgeDAO(db *db.Database) *EdgeDAO {
	return &EdgeDAO{db}
}

func (dao *EdgeDAO) FindAll(directedGraphID uint64) ([]models.Edge, error) {
	// TODO: Add pagination

	dORM := dao.database.GetORM()
	edges := []models.Edge{}
	err := dORM.
		Debug().
		Model(&models.Edge{}).
		Where("directed_graph_id = ?", directedGraphID).
		Limit(100).
		Find(&edges).
		Error
	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (dao *EdgeDAO) FindByPK(directedGraphID, tailNodeID, headNodeID uint64) (*models.Edge, error) {
	dORM := dao.database.GetORM()
	edge := &models.Edge{}
	err := dORM.
		Debug().
		Where("directed_graph_id = ?", directedGraphID).
		Where("tail_node_id = ?", tailNodeID).
		Where("head_node_id = ?", headNodeID).
		Take(edge).
		Error
	if err != nil {
		return nil, err
	}
	return edge, nil
}

func (dao *EdgeDAO) Save(ps *models.Edge) error {
	dORM := dao.database.GetORM()
	err := dORM.Debug().Save(ps).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *EdgeDAO) Delete(directedGraphID, tailNodeID, headNodeID uint64) error {
	dORM := dao.database.GetORM()
	err := dORM.
		Debug().
		Where("directed_graph_id = ?", directedGraphID).
		Where("tail_node_id = ?", tailNodeID).
		Where("head_node_id = ?", headNodeID).
		Delete(&models.Edge{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
