package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./day-3-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var m, n int
	var graph [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		m++
		n = len(line)
		graph = append(graph, []rune{})
		for _, c := range line {
			graph[len(graph)-1] = append(graph[len(graph)-1], c)
		}
	}
	graphCopy := make([][]rune, len(graph))
	for i, row := range graph {
		graphCopy[i] = make([]rune, len(row))
		copy(graphCopy[i], row)
	}
	var explore func(int, int, *[]rune, *[][]int) bool
	explore = func(x int, y int, digits *[]rune, adjIndices *[][]int) bool {
		*digits = append(*digits, graphCopy[x][y])
		graphCopy[x][y] = 'x'
		var deltas [][]int = [][]int{
			{-1, -1},
			{-1, 0},
			{0, -1},
			{1, 0},
			{0, 1},
			{1, 1},
			{-1, 1},
			{1, -1},
		}
		hasAdjacent := false
		for _, delta := range deltas {
			newX := x + delta[0]
			newY := y + delta[1]
			if newX < 0 || newX >= m {
				continue
			}
			if newY < 0 || newY >= n {
				continue
			}
			if graphCopy[newX][newY] == '*' {
				alreadyExists := false
				for _, v := range *adjIndices {
					if newX == v[0] && newY == v[1] {
						alreadyExists = true
					}
				}
				if !alreadyExists {
					*adjIndices = append(*adjIndices, []int{newX, newY})
				}
				hasAdjacent = true
			}
			if graphCopy[newX][newY] >= '0' && graphCopy[newX][newY] <= '9' {
				if explore(newX, newY, digits, adjIndices) {
					hasAdjacent = true
				}
			}
		}
		return hasAdjacent
	}
	var nums int
	var buckets [][][]int = make([][][]int, m)
	for i := range buckets {
		buckets[i] = make([][]int, n)
		for j := range buckets[i] {
			buckets[i][j] = []int{}
		}
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			var digits []rune
			var adjIndices [][]int
			if graphCopy[x][y] >= '0' && graphCopy[x][y] <= '9' && explore(x, y, &digits, &adjIndices) {
				{
					num, err := strconv.Atoi(string(digits))
					for _, adj := range adjIndices {
						buckets[adj[0]][adj[1]] = append(buckets[adj[0]][adj[1]], num)
					}
					if err != nil {
						log.Fatal(err)
					}
					nums++
				}
			}
		}
	}
	var sum int
	for i := range buckets {
		for j := range buckets[i] {
			if len(buckets[i][j]) > 1 {
				currProd := 1
				for _, digit := range buckets[i][j] {
					currProd *= digit
				}
				sum += currProd
			}
		}
	}

	fmt.Println(sum)
}
