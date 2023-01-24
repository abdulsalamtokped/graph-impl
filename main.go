package main

func main() {
	graphImplementation := NewGraph()

	graphImplementation.AddVertex(212)
	graphImplementation.AddVertex(213)
	graphImplementation.AddVertex(214)

	graphImplementation.AddEdge(212, 213)
	graphImplementation.AddEdge(212, 214)
	graphImplementation.AddEdge(213, 214)
	graphImplementation.AddEdge(214, 213)

	graphImplementation.Print()
}
