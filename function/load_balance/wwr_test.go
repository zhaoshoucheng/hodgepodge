package load_balance

import (
	"fmt"
	"testing"
)

func TestWeightRoundRobin_Add(t *testing.T) {
	wrr := WeightRoundRobin{}
	wrr.Add("A","7")
	wrr.Add("B","2")
	wrr.Add("C","1")
	for i := 0; i < 10; i++ {
		fmt.Println(wrr.Next())
	}
}
