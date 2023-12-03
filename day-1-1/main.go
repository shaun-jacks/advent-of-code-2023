package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./day-1-1/input.txt")
	var numbers []int
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var lineCount int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		var firstDigit, lastDigit rune
		var foundFirst bool
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if !foundFirst {
					firstDigit = c
					foundFirst = true
				}
				lastDigit = c
			}
		}
		if foundFirst {
			twoDigitNumberStr := string(firstDigit) + string(lastDigit)
			twoDigitNumber, err := strconv.Atoi(twoDigitNumberStr)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, twoDigitNumber)
		}
	}
	var sum int
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)
}
