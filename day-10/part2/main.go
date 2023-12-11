package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value    string
	Distance int
	I        int
	J        int
	Adj      [][]int
	Loop1    [][]int
	Loop2    [][]int
}

func NewNode(value string, distance, i, j int) *Node {
	return &Node{
		Value:    value,
		Distance: distance,
		I:        i,
		J:        j,
		Adj:      [][]int{},
		Loop1:    [][]int{},
		Loop2:    [][]int{},
	}
}

func MapValueToAdjacencyDiffs(value string) [][]int {
	switch value {
	case "|":
		return [][]int{{-1, 0}, {1, 0}}
	case "-":
		return [][]int{{0, -1}, {0, 1}}
	case "L":
		return [][]int{{-1, 0}, {0, 1}}
	case "J":
		return [][]int{{-1, 0}, {0, -1}}
	case "7":
		return [][]int{{1, 0}, {0, -1}}
	case "F":
		return [][]int{{1, 0}, {0, 1}}
	}
	return [][]int{}
}

func MapValueToLoop1Diffs(value string) [][]int {
	switch value {
	case "|":
		return [][]int{{0, -1}}
	case "-":
		return [][]int{{1, 0}}
	case "L":
		return [][]int{{0, -1}, {1, -1}, {1, 0}}
	case "J":
		return [][]int{{0, 1}, {1, 1}, {1, 0}}
	case "7":
		return [][]int{{-1, 0}, {-1, 1}, {0, 1}}
	case "F":
		return [][]int{{-1, 0}, {-1, -1}, {0, -1}}
	}
	return [][]int{}
}
func MapValueToLoop2Diffs(value string) [][]int {
	switch value {
	case "|":
		return [][]int{{0, 1}}
	case "-":
		return [][]int{{-1, 0}}
	case "L":
		return [][]int{{-1, 1}}
	case "J":
		return [][]int{{-1, -1}}
	case "7":
		return [][]int{{1, -1}}
	case "F":
		return [][]int{{1, 1}}
	}
	return [][]int{}
}
func MapDiffsToIndices(diffs [][]int, i, j, m, n int) [][]int {
	indices := make([][]int, 0)
	for _, diff := range diffs {
		newX, newY := i+diff[0], j+diff[1]
		if 0 <= newX && newX < m && 0 <= newY && newY < n {
			indices = append(indices, []int{newX, newY})
		}
	}
	return indices
}

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Push(node *Node) {
	q.list.PushBack(node)
}

func (q *Queue) Pop() *Node {
	if q.list.Len() == 0 {
		return nil
	}
	front := q.list.Front()
	q.list.Remove(front)
	return front.Value.(*Node)
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
func calculateStartValue(i, j, m, n int, graph [][]string) string {
	var up, down, left, right bool
	for _, dir := range [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	} {
		newX, newY := i+dir[0], j+dir[1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n {
			c := graph[newX][newY]
			if newX == i-1 {
				up = c == "|" || c == "F" || c == "7"
			}
			if newX == i+1 {
				down = c == "|" || c == "J" || c == "L"
			}
			if newY == j-1 {
				left = c == "-" || c == "F" || c == "L"
			}
			if newY == j+1 {
				right = c == "-" || c == "J" || c == "7"
			}
		}
	}
	if up && right {
		return "L"
	}
	if up && left {
		return "J"
	}
	if up && down {
		return "|"
	}
	if down && right {
		return "L"
	}
	if right && left {
		return "-"
	}
	return "-"
}

func main() {
	graph := make([][]string, 0)
	graphNodes := make([][]*Node, 0)
	adjacencyGraph := make([][][]int, 0)
	loopGraph1 := make([][]bool, 0)
	loopGraph2 := make([][]bool, 0)
	mainLoopGraph := make([][]bool, 0)
	f, err := os.Open("./day-10/part2/input.txt")
	if err != nil {
		log.Fatal("Failed to open file", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	startI, startJ := 0, 0
	var lineCount int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0, len(line))
		graphNodeRow := make([]*Node, 0, len(line))
		mainLoopRow := make([]bool, 0, len(line))
		loopGraph1Row := make([]bool, 0, len(line))
		loopGraph2Row := make([]bool, 0, len(line))
		adjacencyRow := make([][]int, 0, len(line))
		fmt.Println(line)
		for idx, c := range line {
			if string(c) == "S" {
				startI, startJ = lineCount, idx
			}
			row = append(row, string(c))
			graphNodeRow = append(graphNodeRow, NewNode(string(c), 0, lineCount, idx))
			mainLoopRow = append(mainLoopRow, false)
			loopGraph1Row = append(loopGraph1Row, false)
			loopGraph2Row = append(loopGraph2Row, false)
			adjacencyRow = append(adjacencyRow, []int{})
		}
		graphNodes = append(graphNodes, graphNodeRow)
		mainLoopGraph = append(mainLoopGraph, mainLoopRow)
		loopGraph1 = append(loopGraph1, loopGraph1Row)
		loopGraph2 = append(loopGraph2, loopGraph2Row)
		adjacencyGraph = append(adjacencyGraph, adjacencyRow)
		graph = append(graph, row)
		lineCount++
	}
	startNode := graphNodes[startI][startJ]
	startNode.Value = calculateStartValue(startI, startJ, len(graph), len(graph[0]), graph)
	for i, row := range graphNodes {
		for j := range row {
			node := graphNodes[i][j]
			adjDiffs, loop1Diffs, loop2Diffs := MapValueToAdjacencyDiffs(node.Value), MapValueToLoop1Diffs(node.Value), MapValueToLoop2Diffs(node.Value)
			node.Adj = MapDiffsToIndices(adjDiffs, node.I, node.J, len(graphNodes), len(row))
			node.Loop1 = MapDiffsToIndices(loop1Diffs, node.I, node.J, len(graphNodes), len(row))
			node.Loop2 = MapDiffsToIndices(loop2Diffs, node.I, node.J, len(graphNodes), len(row))
		}
	}
	q := NewQueue()
	q.Push(startNode)
	maxDist := 0
	for !q.IsEmpty() {
		currNode := q.Pop()
		_, i, j := currNode.Value, currNode.I, currNode.J
		mainLoopGraph[i][j] = true
		loopGraph1[i][j] = false
		loopGraph2[i][j] = false
		graph[currNode.I][currNode.J] = strconv.Itoa(currNode.Distance)
		for _, loop1 := range currNode.Loop1 {
			if !mainLoopGraph[loop1[0]][loop1[1]] {
				loopGraph1[loop1[0]][loop1[1]] = true
			}
		}
		for _, loop2 := range currNode.Loop2 {
			if !mainLoopGraph[loop2[0]][loop2[1]] {
				loopGraph2[loop2[0]][loop2[1]] = true
			}
		}
		for _, adj := range currNode.Adj {
			nextNode := graphNodes[adj[0]][adj[1]]
			if !mainLoopGraph[adj[0]][adj[1]] {
				nextNode.Distance = currNode.Distance + 1
				maxDist = nextNode.Distance
				q.Push(nextNode)
			}
		}
	}
	for _, v := range graph {
		line := make([]string, 0)
		for _, c := range v {
			line = append(line, string(c))
		}
		fmt.Println(strings.Join(line, " "))
	}
	loop1Count := 0
	for _, row := range loopGraph1 {
		line := make([]string, 0)
		for _, v := range row {
			if v {
				loop1Count++
				line = append(line, "1")
			} else {
				line = append(line, ".")
			}
		}
		fmt.Println(strings.Join(line, " "))
	}
	loop2Count := 0
	for _, row := range loopGraph2 {
		line := make([]string, 0)
		for _, v := range row {
			if v {
				loop2Count++
				line = append(line, "2")
			} else {
				line = append(line, ".")
			}
		}
		fmt.Println(strings.Join(line, " "))
	}
	fmt.Println(loop1Count, loop2Count, maxDist)
}
