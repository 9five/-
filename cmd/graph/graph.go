package main

import (
	_graphRepository "algorithm/graph/repository/graph"
	_graphUsecase "algorithm/graph/usecase"
	"fmt"
)

func main() {
	nodes := []string{"A", "B", "C", "D", "E", "F", "G"}
	adjacencyList := make(map[string]map[string]int)
	adjacencyList["A"] = map[string]int{
		"B": 9,
		"C": 2,
	}
	adjacencyList["B"] = map[string]int{
		"A": 9,
		"C": 6,
		"D": 3,
		"E": 1,
	}
	adjacencyList["C"] = map[string]int{
		"A": 2,
		"B": 6,
		"D": 2,
		"F": 9,
	}
	adjacencyList["D"] = map[string]int{
		"B": 3,
		"C": 2,
		"E": 5,
		"F": 6,
	}
	adjacencyList["E"] = map[string]int{
		"B": 1,
		"D": 5,
		"F": 3,
		"G": 7,
	}
	adjacencyList["F"] = map[string]int{
		"C": 9,
		"D": 6,
		"E": 3,
		"G": 4,
	}
	adjacencyList["G"] = map[string]int{
		"E": 7,
		"F": 4,
	}

	start := "A"
	// bellmanFordFunc(nodes, adjacencyList, start)
	dijkstrasFunc(nodes, adjacencyList, start)
}

func bellmanFordFunc(nodes []string, adjacencyList map[string]map[string]int, start string) {
	gRepo := _graphRepository.NewGraphRepository(7)
	if err := gRepo.AddNode(nodes...); err != nil {
		fmt.Println(err.Error())
	}
	if err := gRepo.AddAdjacencyList(adjacencyList, start); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gRepo.Return())
	g := _graphUsecase.NewBellmanFord(gRepo.Return())
	g.Process(start)
	fmt.Println(g.Return())
}

func dijkstrasFunc(nodes []string, adjacencyList map[string]map[string]int, start string) {
	gRepo := _graphRepository.NewGraphRepository(7)
	if err := gRepo.AddNode(nodes...); err != nil {
		fmt.Println(err.Error())
	}
	if err := gRepo.AddAdjacencyList(adjacencyList, start); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(gRepo.Return())
	g := _graphUsecase.NewDjikstras(gRepo.Return())
	g.Process(start)
	fmt.Println(g.Return())
}
