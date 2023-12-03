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
	f, err := os.Open("./day-2-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	gamesMap := make(map[int]map[string]int)
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
		gamesMap[gameId] = map[string]int{}
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
				gamesMap[gameId][rollColor] = rollValue
				if rollValue > maxMap[rollColor] {
					gamesMap[gameId]["exclude"] = 1
				}
			}
		}
	}
	var sum int
	for gameId, v := range gamesMap {
		if _, ok := v["exclude"]; ok {
			if v["exclude"] == 1 {
				continue
			}
		}
		sum += gameId
	}
	fmt.Println(sum)
}
