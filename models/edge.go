package models

type Edge struct {
	DirectedGraphID uint64 `gorm:"primaryKey;autoIncrement:false;index:idx_tail;index:idx_head;required" json:"directed_graph_id" validate:"required"`
	TailNodeID      uint64 `gorm:"primaryKey;autoIncrement:false;index:idx_tail;required" json:"tail_node_id" validate:"required"`
	HeadNodeID      uint64 `gorm:"primaryKey;autoIncrement:false;index:idx_head;required" json:"head_node_id" validate:"required"`
}

func NewEdge(dg *DirectedGraph, tailNode *Node, headNode *Node) *Edge {
	return &Edge{
		DirectedGraphID: dg.ID,
		TailNodeID:      tailNode.ID,
		HeadNodeID:      headNode.ID,
	}
}
