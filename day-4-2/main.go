package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./day-4-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	cards := make([]map[int]int, 1)
	scanner := bufio.NewScanner(f)
	var lines int
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, make(map[int]int))
		lines++
		data := strings.Split(line, " | ")
		inputItem := strings.Split(data[0], ": ")
		cardId, err := strconv.Atoi(strings.Fields(inputItem[0])[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		winningNumbers := strings.Fields(inputItem[1])
		givenNumbers := strings.Fields(data[1])
		for _, winningNumberStr := range winningNumbers {
			winningNumber, err := strconv.Atoi(winningNumberStr)
			if err != nil {
				log.Fatal(err)
				return
			}
			if _, ok := cards[cardId][winningNumber]; !ok {
				cards[cardId][winningNumber] = 0
			}
		}
		for _, givenNumberStr := range givenNumbers {
			givenNumber, err := strconv.Atoi(givenNumberStr)
			if err != nil {
				log.Fatal("error", err)
				return
			}
			if _, ok := cards[cardId][givenNumber]; ok {
				cards[cardId][givenNumber]++
			}
		}
	}

	var answer int
	copies := make([]int, lines+1)
	for i := 0; i <= lines; i++ {
		copies[i] = 1
	}
	for cardId, winningMap := range cards {
		var matches int
		for _, winningMatchCount := range winningMap {
			matches += winningMatchCount
		}
		for i := 0; i < matches; i++ {
			nextCardId := cardId + i + 1
			copies[nextCardId] += copies[cardId]
		}

	}
	for i := 1; i <= lines; i++ {
		answer += copies[i]
	}

	fmt.Println(answer)
}
