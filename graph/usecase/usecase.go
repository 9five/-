package usecase

import (
	"algorithm/domain"
	"container/heap"
	"errors"
)

// 貝爾曼福特演算法
type BellmanFord struct {
	g     domain.Graph
	bs    map[string]domain.BestSoln
	queue []domain.Node
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

		bf.queue = append(bf.queue, bf.g.Nodes[position])
		bf.subexecution(start)
	} else {
		return errors.New("\"" + start + "\" not exist")
	}

	return nil
}

func (bf *BellmanFord) Return() map[string]domain.BestSoln {
	return bf.bs
}

func (bf *BellmanFord) subexecution(start string) {
	if len(bf.queue) == 0 {
		return
	}

	var node domain.Node
	node, bf.queue = bf.queue[0], bf.queue[1:]
	for _, v := range node.Edges {
		if bs, ok := bf.bs[v.To]; !ok || bs.Weight > bf.bs[node.Sign].Weight+v.Weight {
			bf.bs[v.To] = domain.BestSoln{PrevSign: node.Sign, Weight: bf.bs[node.Sign].Weight + v.Weight}
			_, position := bf.g.Contain(v.To)
			bf.queue = append(bf.queue, bf.g.Nodes[position])
		}
	}
	bf.subexecution(start)
}

// 戴克斯特拉演算法
type Dijkstras struct {
	g  domain.Graph
	pq domain.PriorityQueue
	bs map[string]domain.BestSoln
}

func NewDjikstras(g domain.Graph) domain.GraphAlgorithmUsecase {
	return &Dijkstras{
		g:  g,
		bs: make(map[string]domain.BestSoln),
	}
}

func (d *Dijkstras) Process(start string) error {
	if exist, _ := d.g.Contain(start); exist {
		d.bs[start] = domain.BestSoln{Weight: 0}
		d.pq = append(d.pq, &domain.Item{
			Value:    start,
			Priority: 0,
			Index:    0,
		})
		heap.Init(&d.pq)

		d.subexecution()
	} else {
		return errors.New("\"" + start + "\" not exist")
	}

	return nil
}

func (d *Dijkstras) subexecution() {
	if d.pq.Len() <= 0 {
		return
	}

	item := heap.Pop(&d.pq).(*domain.Item)
	_, position := d.g.Contain(item.Value)
	node := d.g.Nodes[position]

	for _, v := range node.Edges {
		if bs, ok := d.bs[v.To]; !ok || bs.Weight > d.bs[node.Sign].Weight+v.Weight {
			d.bs[v.To] = domain.BestSoln{PrevSign: node.Sign, Weight: d.bs[node.Sign].Weight + v.Weight}
			heap.Push(&d.pq, &domain.Item{
				Value:    v.To,
				Priority: d.bs[v.To].Weight,
			})
		}
	}
	d.subexecution()
}

func (d *Dijkstras) Return() map[string]domain.BestSoln {
	return d.bs
}
