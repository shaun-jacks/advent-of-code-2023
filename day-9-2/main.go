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
	ParentLeft  *Node
	ParentRight *Node
	Value       int
	Level       int
	ChildLeft   *Node
	ChildRight  *Node
}

func NewNode(value, level int) *Node {
	return &Node{
		Value: value,
		Level: level,
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

func (q *Queue) printQueue() {
	if q.list.Len() == 0 {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Println("Queue contents:")
	for element := q.list.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, " ")
	}
	fmt.Println()
}

func processLine(line string, q *Queue) (int, error) {
	rootNodeValues := strings.Split(line, " ")
	for _, v := range rootNodeValues {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Failed to convert to int", err)
			return 0, err
		}
		node := NewNode(value, 0)
		q.Push(node)
	}
	var prev, curr, head *Node
	prev = q.Pop()
	head = prev
	allZerosInLevel := true
	for !q.IsEmpty() {
		curr = q.Pop()
		if curr.Level != prev.Level {
			head = curr
			prev = curr
			if allZerosInLevel {
				for !q.IsEmpty() {
					q.Pop()
				}
				break
			}
			allZerosInLevel = true
			continue
		}
		newValue := curr.Value - prev.Value
		if newValue != 0 {
			allZerosInLevel = false
		}
		childNode := NewNode(newValue, curr.Level+1)
		childNode.ParentLeft = prev
		childNode.ParentRight = curr
		q.Push(childNode)
		prev = curr
	}
	newHead := NewNode(0, curr.Level)
	newHead.ParentRight = head.ParentLeft
	curr = newHead
	for curr.ParentRight != nil {
		parentLeftNode := NewNode(curr.ParentRight.Value-curr.Value, curr.Level-1)
		parentLeftNode.ParentRight = curr.ParentRight.ParentLeft
		curr.ParentLeft = parentLeftNode
		curr = parentLeftNode
	}
	return curr.Value, nil
}

func main() {
	q := NewQueue()
	f, err := os.Open("./day-9-2/input.txt")
	if err != nil {
		log.Fatal("Failed to open file", err)
		return
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var answer int
	for scanner.Scan() {
		line := scanner.Text()
		nextVal, err := processLine(line, q)
		if err != nil {
			log.Fatal("Failed to process line", err)
			return
		}
		answer += nextVal
	}
	fmt.Println(answer)
}
