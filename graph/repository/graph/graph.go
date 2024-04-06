package repository

import (
	"algorithm/domain"
	"errors"
	"fmt"
)

type GraphRepository struct {
	g domain.Graph
}

func NewGraphRepository(nn int) domain.GraphRepository {
	return &GraphRepository{domain.Graph{NodeNum: nn}}
}

func (gr *GraphRepository) AddNode(s ...string) error {
	if len(gr.g.Nodes)+len(s) > gr.g.NodeNum {
		return errors.New("NODES OUT OF STACK!")
	}
	for _, v := range s {
		if exist, _ := gr.g.Contain(v); exist {
			fmt.Printf("\"%s\" is already signed up.\n", v)
		}
		gr.g.Nodes = append(gr.g.Nodes, domain.Node{Sign: v})
	}

	return nil
}

func (gr *GraphRepository) AddEdge(from, to string, weight int) error {
	if exist, _ := gr.g.Contain(to); !exist {
		return errors.New("\"" + to + "\" not exist 1")
	}

	if exist, position := gr.g.Contain(from); !exist {
		return errors.New("\"" + from + "\" not exist 2")
	} else {
		gr.g.Nodes[position].Edges = append(gr.g.Nodes[position].Edges, domain.Edge{From: from, To: to, Weight: weight})
	}

	return nil
}

func (gr *GraphRepository) AddAdjacencyList(adjacencyList map[string]map[string]int, start string) error {
	for k1, v1 := range adjacencyList {
		for k2, v2 := range v1 {
			if k2 == start {
				continue
			}
			if err := gr.AddEdge(k1, k2, v2); err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	return nil
}

func (gr *GraphRepository) Return() domain.Graph {
	return gr.g
}
