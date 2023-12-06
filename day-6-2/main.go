package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputFile(filepath string) ([]int, []int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return []int{}, []int{}, err
	}
	defer f.Close()
	var distanceComb string
	var timeComb string
	var times []int
	var distances []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			timeStrings := strings.Fields(line[5:])
			for _, timeString := range timeStrings {
				timeComb += timeString
			}
		}
		if strings.HasPrefix(line, "Distance:") {
			distanceStrings := strings.Fields(line[9:])
			for _, distanceString := range distanceStrings {
				distanceComb += distanceString
			}
		}
	}
	time, err := strconv.Atoi(timeComb)
	if err != nil {
		log.Fatal(err)
		return []int{}, []int{}, err
	}
	distance, err := strconv.Atoi(distanceComb)
	if err != nil {
		log.Fatal(err)
		return []int{}, []int{}, err
	}
	times = append(times, time)
	distances = append(distances, distance)
	return times, distances, nil
}

func main() {
	times, distances, err := readInputFile("./day-6-1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(times, distances)
	raceDistances := make([]int, 0, len(distances))
	runningProduct := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		var betterTimes int
		for speed := 0; speed <= time; speed++ {
			timeLeft := time - speed
			raceDistance := timeLeft * speed
			if raceDistance > distance {
				betterTimes++
			}
		}
		raceDistances = append(raceDistances, betterTimes)
		runningProduct *= betterTimes
	}
	fmt.Println(raceDistances)
	fmt.Println(runningProduct)
}
