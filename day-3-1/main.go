package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./day-3-1/input.txt")
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
	var explore func(int, int, *[]rune) bool
	explore = func(x int, y int, digits *[]rune) bool {
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
			if graphCopy[newX][newY] != '.' &&
				graphCopy[newX][newY] != 'x' &&
				!(graphCopy[newX][newY] >= '0' && graphCopy[newX][newY] <= '9') {
				hasAdjacent = true
			}
			if graphCopy[newX][newY] >= '0' && graphCopy[newX][newY] <= '9' {
				if explore(newX, newY, digits) {
					hasAdjacent = true
				}
			}
		}
		return hasAdjacent
	}
	var nums int
	var sum int
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			var digits []rune
			if graphCopy[x][y] >= '0' && graphCopy[x][y] <= '9' && explore(x, y, &digits) {
				{
					num, err := strconv.Atoi(string(digits))
					if err != nil {
						log.Fatal(err)
					}
					sum += num
					nums++

				}
			}
		}
	}

	fmt.Println(sum)
}
