package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./day-4-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	cardMap := make(map[int]map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " | ")
		inputItem := strings.Split(data[0], ": ")
		cardId, err := strconv.Atoi(strings.Fields(inputItem[0])[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		winningNumbers := strings.Fields(inputItem[1])
		givenNumbers := strings.Fields(data[1])
		if _, ok := cardMap[cardId]; !ok {
			cardMap[cardId] = make(map[int]int)
		}
		for _, winningNumberStr := range winningNumbers {
			winningNumber, err := strconv.Atoi(winningNumberStr)
			if err != nil {
				log.Fatal(err)
				return
			}
			cardMap[cardId][winningNumber] = 0
		}
		for _, givenNumberStr := range givenNumbers {
			givenNumber, err := strconv.Atoi(givenNumberStr)
			if err != nil {
				log.Fatal("error", err)
				return
			}
			if _, ok := cardMap[cardId][givenNumber]; ok {
				cardMap[cardId][givenNumber]++
			}
		}
	}

	var answer int
	for _, winningMap := range cardMap {
		var matches int
		for _, winningMatchCount := range winningMap {
			matches += winningMatchCount
		}
		if matches > 0 {
			res := math.Pow(2, float64(matches)-1)
			answer += int(res)
		}

	}

	fmt.Println(answer)
}
