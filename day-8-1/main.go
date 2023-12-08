package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func readInputFile(filepath string) (map[string]*Node, []string, string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return map[string]*Node{}, []string{}, "", err
	}
	defer f.Close()
	nodes := make(map[string]*Node)
	steps := make([]string, 0)
	var rootNode string = ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "=") == -1 && (strings.HasPrefix(line, "L") || strings.HasPrefix(line, "R")) {
			for _, c := range line {
				steps = append(steps, string(c))
			}
			continue
		}
		if strings.Index(line, "=") == -1 {
			continue
		}
		parts := strings.Split(line, "=")
		value := strings.TrimSpace(parts[0])
		connections := strings.Trim(strings.TrimSpace(parts[1]), "()")
		connectionParts := strings.Split(connections, ",")
		leftValue := strings.TrimSpace(connectionParts[0])
		rightValue := strings.TrimSpace(connectionParts[1])
		if rootNode == "" {
			rootNode = value
		}
		if _, exists := nodes[value]; !exists {
			nodes[value] = &Node{Value: value}
		}
		if _, exists := nodes[leftValue]; !exists {
			nodes[leftValue] = &Node{Value: leftValue}
		}
		if _, exists := nodes[rightValue]; !exists {
			nodes[rightValue] = &Node{Value: rightValue}
		}
		nodes[value].Left = nodes[leftValue]
		nodes[value].Right = nodes[rightValue]
	}
	return nodes, steps, rootNode, nil
}

func stepsRequired(nodes map[string]*Node, root string, steps []string) int {
	node := nodes[root]
	if node.Value == "ZZZ" {
		return 0
	}
	var stepCount int
	index := 0
	for node.Value != "ZZZ" {
		step := steps[index]
		fmt.Println(node.Value)
		if step == "L" {
			node = node.Left
		}
		if step == "R" {
			node = node.Right
		}
		stepCount++
		fmt.Println(index, steps[index], stepCount, node.Value)
		index = (index + 1) % len(steps)
	}
	return stepCount
}

func main() {
	nodes, steps, rootNode, err := readInputFile("./day-8-1/input.txt")
	fmt.Println(steps)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(nodes, rootNode, steps, stepsRequired(nodes, "AAA", steps))
	return
}
