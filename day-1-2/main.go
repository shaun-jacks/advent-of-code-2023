package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stringNums []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	f, err := os.Open("./day-1-2/input.txt")
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
		var firstDigitLocation, lastDigitLocation int
		var foundFirst bool
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if !foundFirst {
					firstDigit = c
					firstDigitLocation = i
					foundFirst = true
				}
				lastDigit = c
				lastDigitLocation = i
			}
		}
		for i, num := range stringNums {
			idx := strings.Index(line, num)
			if idx == -1 {
				continue
			}
			c := rune(i + '1')
			if !foundFirst || (foundFirst && idx < firstDigitLocation) {
				firstDigitLocation = idx
				firstDigit = c
				foundFirst = true
			}
			lastIdx := strings.LastIndex(line, num)
			if lastIdx > lastDigitLocation {
				lastDigitLocation = lastIdx
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
