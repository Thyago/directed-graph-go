package models

type DirectedGraph struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" validate:"required" json:"name"`
}

func NewDirectedGraph(name string) *DirectedGraph {
	return &DirectedGraph{
		Name: name,
	}
}
