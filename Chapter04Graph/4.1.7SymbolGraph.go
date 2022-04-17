package Chapter04Graph

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type SymbolGraph struct {
	// 键为图的顶点，值为顶点的整数表示
	// 值为0时，表示顶点不存在，从1开始计算。
	index map[string]int
	graph *Graph
}

func NewSymbolGraph(filePath string) *SymbolGraph {
	sg := new(SymbolGraph)
	sg.index = make(map[string]int)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open error")
		return nil
	}
	defer file.Close()
	count := 1
	br := bufio.NewReader(file)
	for {
		row, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data := string(row)
		splits := strings.Split(data, " ")
		for _, v := range splits {
			if sg.index[v] == 0 {
				sg.index[v] = count
				count++
			}
		}
	}
	sg.graph = NewGraph(count-1)
	br = bufio.NewReader(file)
	for {
		row, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data := string(row)
		splits := strings.Split(data, " ")
		v := sg.index[splits[0]] - 1
		w := sg.index[splits[1]] - 1
		sg.graph.AddEdge(v, w)
	}
	return sg
}

func (sg *SymbolGraph) Contains(s string) bool {
	return sg.index[s] != 0
}

func (sg *SymbolGraph) Index(s string) int {
	return sg.index[s] - 1
}

func (sg *SymbolGraph) Name(v int) string {
	for key, val := range sg.index {
		if v == val - 1 {
			return key
		}
	}
	return ""
}