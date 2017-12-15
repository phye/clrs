package graph

import (
	"errors"
	"fmt"
)

type (
	MGraph struct {
		matrix   [][]int
		directed bool
		values   []interface{}
		indices  map[interface{}]int // Interface to Int for quick lookup
	}
)

func NewMatrixGraph(directed bool) *MGraph {
	mg := &MGraph{
		make([][]int, 0),
		directed,
		make([]interface{}, 0),
		make(map[interface{}]int),
	}
	return mg
}

// Do note that v must be comparable
func (mg *MGraph) AddNode(v interface{}) {
	ol := len(mg.matrix)
	for i := 0; i < ol; i++ {
		mg.matrix[i] = append(mg.matrix[i], 0)
	}
	mg.matrix = append(mg.matrix, make([]int, ol+1))
	mg.values = append(mg.values, v)
	mg.indices[v] = ol
}

func (mg *MGraph) AddEdge(sv, ev interface{}, weight int) error {
	if err := mg.checkExist(sv); err != nil {
		return err
	}
	if err := mg.checkExist(ev); err != nil {
		return err
	}
	i, j := mg.indices[sv], mg.indices[ev]
	mg.matrix[i][j] = weight
	if !mg.directed {
		mg.matrix[j][i] = weight
	}
	return nil
}

func (mg *MGraph) String() string {
	ret := fmt.Sprintf("========== Adjacent Matrix Graph ==========\n")
	for i := 0; i < len(mg.matrix); i++ {
		if i == 0 {
			ret += fmt.Sprintf("%8v ", "")
			for j := 0; j < len(mg.values); j++ {
				ret += fmt.Sprintf("%8v ", mg.values[j])
			}
			ret += fmt.Sprintf("\n")
		}
		for j := 0; j < len(mg.matrix[i]); j++ {
			if j == 0 {
				ret += fmt.Sprintf("%8v ", mg.values[i])
			}
			ret += fmt.Sprintf("%8v ", mg.matrix[i][j])
		}
		ret += fmt.Sprintf("\n")
	}
	ret += fmt.Sprintf("===========================================\n")
	return ret
}

func (mg *MGraph) Adj(tv interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	if err := mg.checkExist(tv); err != nil {
		return ret, err
	}
	i := mg.indices[tv]
	for j := 0; j < len(mg.matrix[i]); j++ {
		if mg.matrix[i][j] > 0 {
			ret = append(ret, mg.values[j])
		}
	}
	return ret, nil
}

func (mg *MGraph) checkExist(tv interface{}) error {
	if _, ok := mg.indices[tv]; !ok {
		msg := fmt.Sprintf("ERROR: %v does not exist in graph", tv)
		fmt.Println(msg)
		return errors.New(msg)
	}
	return nil
}

func (mg *MGraph) Weight(sv, ev interface{}) (int, error) {
	if err := mg.checkExist(sv); err != nil {
		return -1, err
	}
	if err := mg.checkExist(ev); err != nil {
		return -1, err
	}
	i, j := mg.indices[sv], mg.indices[ev]
	return mg.matrix[i][j], nil
}
