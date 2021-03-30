package models

type Node struct {
	ID              uint64 `gorm:"primaryKey;index:idx_node;autoIncrement" json:"id"`
	DirectedGraphID uint64 `gorm:"required;index:idx_node;not null" json:"directed_graph_id"`
	DirectedGraph   *DirectedGraph
	NodeMetadata    []NodeMetadata `json:"metadata"`
}

type NodeMetadata struct {
	NodeID  uint64 `gorm:"primaryKey;required;not null;index" json:"node_id"`
	Meta    string `gorm:"primaryKey;required;not null" json:"meta"`
	Content string `json:"content"`
}

func NewNode(dg *DirectedGraph, metadata []NodeMetadata) *Node {
	return &Node{
		DirectedGraph:   dg,
		DirectedGraphID: dg.ID,
		NodeMetadata:    metadata,
	}
}
