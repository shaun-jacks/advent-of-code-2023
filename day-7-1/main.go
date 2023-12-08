package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var rankingMap map[string]int = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Draw struct {
	Hand string
	Bid  int
}

func readInputFile(filepath string) ([]Draw, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return []Draw{}, err
	}
	defer f.Close()
	draws := make([]Draw, 0)
	hands := make([]string, 0)
	bids := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		hand := parts[0]
		hands = append(hands, hand)
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Failed to read bid", err)
			return []Draw{}, err
		}
		bids = append(bids, bid)
		draws = append(draws, Draw{
			Bid:  bid,
			Hand: hand,
		})
	}
	return draws, nil
}

func rankHand(draw Draw) int {
	repeatMap := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	charCount := make(map[string]int)
	for _, c := range draw.Hand {
		if _, ok := charCount[string(c)]; !ok {
			charCount[string(c)] = 1
		} else {
			charCount[string(c)]++
		}
	}
	for _, v := range charCount {
		repeatMap[v]++
	}

	if repeatMap[5] == 1 {
		return 6
	}
	if repeatMap[4] == 1 {
		return 5
	}
	if repeatMap[3] == 1 && repeatMap[2] == 1 {
		return 4
	}
	if repeatMap[3] == 1 && repeatMap[1] == 2 {
		return 3
	}
	if repeatMap[2] == 2 && repeatMap[1] == 1 {
		return 2
	}
	if repeatMap[2] == 1 && repeatMap[1] == 3 {
		return 1
	}
	if repeatMap[1] == 5 {
		return 0
	}
	return 0
}

func main() {
	draws, err := readInputFile("./day-7-1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(draws)
	rankBuckets := make([][]Draw, 8)
	for _, draw := range draws {
		rank := rankHand(draw)
		rankBuckets[rank] = append(rankBuckets[rank], draw)
	}
	for i, bucketedDraws := range rankBuckets {
		sort.Slice(bucketedDraws, func(i, j int) bool {
			for k := 0; k < len(bucketedDraws[i].Hand); k++ {
				handA, handB := rankingMap[string(bucketedDraws[i].Hand[k])], rankingMap[string(bucketedDraws[j].Hand[k])]
				if handA != handB {
					return handA < handB
				}
			}
			return true
		})
		rankBuckets[i] = bucketedDraws
	}
	orderedRankedDraws := make([]Draw, 0, len(draws))
	for i := range rankBuckets {
		for j := range rankBuckets[i] {
			orderedRankedDraws = append(orderedRankedDraws, rankBuckets[i][j])
		}
	}
	fmt.Println(orderedRankedDraws)
	var answer int
	for i := range orderedRankedDraws {
		rank := i + 1
		answer += (rank * orderedRankedDraws[i].Bid)
	}
	fmt.Println(answer)
}
