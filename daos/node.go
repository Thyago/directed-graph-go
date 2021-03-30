package daos

import (
	"github.com/thyago/directed-graph-go/db"
	"github.com/thyago/directed-graph-go/models"
)

type NodeDAO struct {
	database *db.Database
}

func NewNodeDAO(db *db.Database) *NodeDAO {
	return &NodeDAO{db}
}

func (dao *NodeDAO) FindAll(directedGraphID uint64) ([]models.Node, error) {
	// TODO: Add pagination

	dORM := dao.database.GetORM()
	nodes := []models.Node{}
	err := dORM.
		Debug().
		Preload("NodeMetadata").
		Model(&models.Node{}).
		Where("directed_graph_id = ?", directedGraphID).
		Limit(100).
		Find(&nodes).
		Error
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (dao *NodeDAO) FindByID(directedGraphID, id uint64) (*models.Node, error) {
	dORM := dao.database.GetORM()
	node := &models.Node{}
	err := dORM.
		Debug().
		Preload("NodeMetadata").
		Where("directed_graph_id = ?", directedGraphID).
		Where("id = ?", id).
		Take(node).
		Error
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (dao *NodeDAO) Save(n *models.Node) error {
	dORM := dao.database.GetORM()
	if n.ID != 0 {
		// Remove old metadata
		err := dORM.
			Debug().
			Where("node_id = ?", n.ID).
			Delete(&models.NodeMetadata{}).
			Error
		if err != nil {
			return err
		}
	}
	err := dORM.Debug().Save(n).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *NodeDAO) Delete(directedGraphID, id uint64) error {
	dORM := dao.database.GetORM()
	err := dORM.
		Debug().
		Where("directed_graph_id = ?", directedGraphID).
		Where("id = ?", id).
		Delete(&models.Node{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
