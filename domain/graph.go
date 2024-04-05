package domain

import ()

type Graph struct {
	NodeNum int
	Nodes   []Node
}

type Node struct { // somewhere will call it vertex
	Sign  string
	Edges []Edge
}

type Edge struct {
	From   string
	To     string
	Weight int
}

type BestSoln struct {
	// Sign     string
	PrevSign string
	Weight   int
}

type GraphRepository interface {
	AddNode(s ...string) error
	AddEdge(form, to string, weight int) error
	Return() Graph
}

type GraphAlgorithmUsecase interface {
	Process(start string) error
	Return() map[string]BestSoln
}

func (g *Graph) Contain(s string) (bool, int) {
	for k, v := range g.Nodes {
		if v.Sign == s {
			return true, k
		}
	}
	return false, 0
}
