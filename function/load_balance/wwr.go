package load_balance

import (
	"errors"
	"strconv"
)

/*负载均衡算法
wwr 加权轮询算法
（Weighted Round Robin）
*/

type Element struct {
	Data   string
	weight int
}

type IWrr interface {
	GetNext() string
}

//版本1 随机数版本
//将整体权重求和，计算每一个权重的比例，根据比例获取随机数，当调用量大时比较适用

//版本2 Nginx版本，既可以随机，又可以平滑负载
type WeightRoundRobin struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}
type WeightNode struct {
	addr            string
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

func (r *WeightRoundRobin) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobin) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		/*
		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
		 */
		total += w.weight
		w.currentWeight += w.weight
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	best.currentWeight -= total
	return best.addr
}


