package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxMap map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	f, err := os.Open("./day-2-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	gamesMap := make(map[int]map[string]map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		game := parts[0]
		gameIdStr := strings.Split(game, " ")[1]
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			log.Fatal(err)
			return
		}
		gamesMap[gameId] = map[string]map[string]int{

			"max": {
				"include": 0,
			},
		}
		gameRounds := strings.Split(parts[1], "; ")
		for _, gameRound := range gameRounds {
			rolls := strings.Split(gameRound, ", ")
			for _, v := range rolls {
				roll := strings.Split(v, " ")
				rollValue, err := strconv.Atoi(roll[0])
				if err != nil {
					log.Fatal(err)
					return
				}
				rollColor := roll[1]
				if _, ok := gamesMap[gameId][rollColor]; !ok {
					gamesMap[gameId][rollColor] = map[string]int{}
					gamesMap[gameId][rollColor]["max"] = rollValue
					gamesMap[gameId][rollColor]["min"] = rollValue
				}
				if rollValue > gamesMap[gameId][rollColor]["max"] {
					gamesMap[gameId][rollColor]["max"] = rollValue
				}
				if rollValue < gamesMap[gameId][rollColor]["min"] {
					gamesMap[gameId][rollColor]["min"] = rollValue
				}
				if rollValue > maxMap[rollColor] {
					gamesMap[gameId]["max"]["include"] = 1
				}
			}
		}
	}
	var sum int
	for _, v := range gamesMap {
		var power int = 1
		power *= v["green"]["max"]
		power *= v["red"]["max"]
		power *= v["blue"]["max"]
		sum += power
	}
	fmt.Println(sum)
}
