package usecase

import (
	"algorithm/domain"
	"errors"
	"fmt"
)

// 貝爾曼福特演算法
type BellmanFord struct {
	g     domain.Graph
	bs    map[string]domain.BestSoln
	quene []domain.Node
}

func NewBellmanFord(g domain.Graph) domain.GraphAlgorithmUsecase {
	return &BellmanFord{
		g:  g,
		bs: make(map[string]domain.BestSoln),
	}
}

func (bf *BellmanFord) Process(start string) error {
	if exist, position := bf.g.Contain(start); exist {
		bf.bs[start] = domain.BestSoln{Weight: 0}

		bf.quene = append(bf.quene, bf.g.Nodes[position])
		bf.subexecution(start)
	} else {
		return errors.New("\"" + start + "\" not exist")
	}
	fmt.Println()

	return nil
}

func (bf *BellmanFord) Return() map[string]domain.BestSoln {
	return bf.bs
}

func (bf *BellmanFord) subexecution(start string) {
	if len(bf.quene) == 0 {
		return
	}
	var node domain.Node
	node, bf.quene = bf.quene[0], bf.quene[1:]
	for _, v := range node.Edges {
		if bs, ok := bf.bs[v.To]; !ok || bs.Weight > bf.bs[node.Sign].Weight+v.Weight {
			bf.bs[v.To] = domain.BestSoln{PrevSign: node.Sign, Weight: bf.bs[node.Sign].Weight + v.Weight}
			_, position := bf.g.Contain(v.To)
			bf.quene = append(bf.quene, bf.g.Nodes[position])
		}
	}
	bf.subexecution(start)
}

// 戴克斯特拉演算法
type Dijkstras struct {
	g  domain.Graph
	bs map[string]domain.BestSoln
}

func NewDjikstras(g domain.Graph) domain.GraphAlgorithmUsecase {
	return &Dijkstras{
		g:  g,
		bs: make(map[string]domain.BestSoln),
	}
}

func (d *Dijkstras) Process(start string) error {
	return nil
}

func (d *Dijkstras) Return() map[string]domain.BestSoln {
	return d.bs
}
