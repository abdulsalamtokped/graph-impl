package main

import "fmt"

type (
	Graph struct {
		vertices []*Vertex
	}
	Vertex struct {
		key      int
		adjacent []*Vertex
	}

	GraphInterface interface {
		AddVertex(vertex int) error
		AddEdge(to int, from int) error
		Print()
	}
)

func NewGraph() GraphInterface {
	return &Graph{}
}
func (g *Graph) AddVertex(vertex int) error {
	if contains(g.vertices, vertex) {
		err := fmt.Errorf("Vertex %d already exists", vertex)
		return err
	} else {
		v := &Vertex{
			key: vertex,
		}
		g.vertices = append(g.vertices, v)
	}
	return nil
}
func (g *Graph) AddEdge(to, from int) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)

	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("Not a valid edge from %d ---> %d", from, to)
	}
	if contains(fromVertex.adjacent, toVertex.key) {
		return fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.key, toVertex.key)
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	return nil
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%d ", v.key)
		}
		fmt.Println()
	}
}

// helper
// getVertex will return a vertex point if exists or return nil
func (g *Graph) getVertex(vertex int) *Vertex {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}
func contains(v []*Vertex, key int) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}
	return false
}
