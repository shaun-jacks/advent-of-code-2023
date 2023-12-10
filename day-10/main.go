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
}

func NewNode(value string, distance, i, j int) *Node {
	return &Node{
		Value:    value,
		Distance: distance,
		I:        i,
		J:        j,
	}
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
	visited := make([][]bool, 0)
	f, err := os.Open("./day-10/input.txt")
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
		visitedRow := make([]bool, 0, len(line))
		fmt.Println(line)
		for idx, c := range line {
			if string(c) == "S" {
				startI, startJ = lineCount, idx
			}
			row = append(row, string(c))
			visitedRow = append(visitedRow, false)
		}
		graph = append(graph, row)
		visited = append(visited, visitedRow)
		lineCount++
	}

	startNode := NewNode(calculateStartValue(startI, startJ, len(graph), len(graph[0]), graph), 0, startI, startJ)
	q := NewQueue()
	q.Push(startNode)
	maxDist := 0
	for !q.IsEmpty() {
		currNode := q.Pop()
		c, i, j := currNode.Value, currNode.I, currNode.J
		visited[currNode.I][currNode.J] = true
		graph[currNode.I][currNode.J] = strconv.Itoa(currNode.Distance)
		fromUp := c == "|" || c == "L" || c == "J" && i-1 >= 0 && i-1 < len(graph)
		fromDown := c == "|" || c == "7" || c == "F" && i+1 >= 0 && i+1 < len(graph)
		fromLeft := c == "-" || c == "7" || c == "J" && j-1 >= 0 && j-1 < len(graph[0])
		fromRight := c == "-" || c == "L" || c == "F" && j+1 >= 0 && j+1 < len(graph[0])
		if fromUp && !visited[i-1][j] {
			maxDist = currNode.Distance + 1
			q.Push(NewNode(graph[i-1][j], currNode.Distance+1, i-1, j))
		}
		if fromDown && !visited[i+1][j] {
			maxDist = currNode.Distance + 1
			q.Push(NewNode(graph[i+1][j], currNode.Distance+1, i+1, j))
		}
		if fromLeft && !visited[i][j-1] {
			maxDist = currNode.Distance + 1
			q.Push(NewNode(graph[i][j-1], currNode.Distance+1, i, j-1))
		}
		if fromRight && !visited[i][j+1] {
			maxDist = currNode.Distance + 1
			q.Push(NewNode(graph[i][j+1], currNode.Distance+1, i, j+1))
		}
	}
	for _, v := range graph {
		line := make([]string, 0)
		for _, c := range v {
			line = append(line, string(c))
		}
		fmt.Println(strings.Join(line, " "))
	}
	fmt.Println(maxDist)
}
