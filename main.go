package main

func main() {
	graphImplementation := NewGraph()

	graphImplementation.AddVertex(1)
	graphImplementation.AddVertex(11)
	graphImplementation.AddVertex(34)

	graphImplementation.AddEdge(1, 11)
	graphImplementation.AddEdge(11, 34)
	graphImplementation.AddEdge(34, 44)

	graphImplementation.Print()
}
