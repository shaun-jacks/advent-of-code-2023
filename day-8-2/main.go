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

func readInputFile(filepath string) (map[string]*Node, []string, []string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return map[string]*Node{}, []string{}, []string{}, err
	}
	defer f.Close()
	nodes := make(map[string]*Node)
	steps := make([]string, 0)
	rootNodes := make([]string, 0)
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
		if _, exists := nodes[value]; !exists {
			nodes[value] = &Node{Value: value}
			if string(value[2]) == "A" {
				rootNodes = append(rootNodes, value)
			}
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
	return nodes, steps, rootNodes, nil
}

func allEndWithZ(nodes []*Node) bool {
	allEndWithZ := true
	for _, node := range nodes {
		if string(node.Value[2]) != "Z" {
			allEndWithZ = false
		}
	}
	return allEndWithZ
}

func stepsRequired(nodes map[string]*Node, rootNode string, steps []string) int {
	node := nodes[rootNode]
	if string(node.Value[2]) == "Z" {
		return 0
	}
	var stepCount int
	index := 0
	for string(node.Value[2]) != "Z" {
		step := steps[index]
		if step == "L" {
			node = node.Left
		}
		if step == "R" {
			node = node.Right
		}
		stepCount++
		index = (index + 1) % len(steps)
	}
	return stepCount
}

// gcd computes the Greatest Common Divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm computes the Least Common Multiple of a and b.
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// lcmSlice computes the LCM of a slice of integers.
func lcmSlice(nums []int) int {
	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}

func main() {
	nodes, steps, rootNodes, err := readInputFile("./day-8-2/input.txt")
	fmt.Println(steps)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	allSteps := make([]int, 0)
	for _, rootNode := range rootNodes {
		stepsNeeded := stepsRequired(nodes, rootNode, steps)
		allSteps = append(allSteps, stepsNeeded)
	}

	fmt.Println(lcmSlice(allSteps))
	return
}
