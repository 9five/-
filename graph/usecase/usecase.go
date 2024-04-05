package usecase

import (
	"algorithm/domain"
	"errors"
)

// 貝爾曼福特演算法
type BellmanFord struct {
	g  domain.Graph
	bs map[string]domain.BestSoln
}

func NewBellmanFord(g domain.Graph) domain.GraphAlgorithmUsecase {
	return &BellmanFord{
		g:  g,
		bs: make(map[string]domain.BestSoln),
	}
}

func (bf *BellmanFord) Process(start string) error {
	if exist, _ := bf.g.Contain(start); !exist {
		return errors.New("\"" + start + "\" not exist")
	}

	bf.bs[start] = domain.BestSoln{Weight: 0}

	_, position := bf.g.Contain(start)
	bf.subexecution(0, bf.g.Nodes[position])

	return nil
}

func (bf *BellmanFord) Return() map[string]domain.BestSoln {
	return bf.bs
}

func (bf *BellmanFord) subexecution(execNum int, node domain.Node) {
	if execNum >= len(bf.g.Nodes) {
		return
	}
	for _, v := range node.Edges {
		if bs, ok := bf.bs[v.To]; !ok || bs.Weight > bf.bs[node.Sign].Weight+v.Weight {
			bf.bs[v.To] = domain.BestSoln{PrevSign: node.Sign, Weight: bf.bs[node.Sign].Weight + v.Weight}
		}
		_, position := bf.g.Contain(v.To)
		bf.subexecution(execNum+1, bf.g.Nodes[position])
	}
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
