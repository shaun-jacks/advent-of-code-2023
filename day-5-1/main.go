package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	Source      int
	Destination int
	Increment   int
}

type Mapping struct {
	Edges []Edge
}

func readMappingsFromFile(filename string) ([]Mapping, []int, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var mappings []Mapping
	var seeds []int
	var currentMapping *Mapping

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seedStrs := strings.Split(line[7:], " ")
			for _, s := range seedStrs {
				seed, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, fmt.Errorf("invalid seed value: %s", s)
				}
				seeds = append(seeds, seed)
			}
			continue
		}

		if strings.Contains(line, "map:") {
			if currentMapping != nil {
				mappings = append(mappings, *currentMapping)
			}
			currentMapping = &Mapping{}
			continue
		}

		if currentMapping != nil {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				source, err1 := strconv.Atoi(parts[1])
				destination, err2 := strconv.Atoi(parts[0])
				increment, err3 := strconv.Atoi(parts[2])
				if err1 != nil || err2 != nil || err3 != nil {
					return nil, nil, fmt.Errorf("invalid edge data: %s", line)
				}
				currentMapping.Edges = append(currentMapping.Edges, Edge{
					Source:      source,
					Destination: destination,
					Increment:   increment,
				})
			}
		}
	}
	// Add the last mapping if exists
	if currentMapping != nil && len(currentMapping.Edges) > 0 {
		mappings = append(mappings, *currentMapping)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return mappings, seeds, nil
}

func getLocationFromSeed(mappings []Mapping, seed int) int {
	var start int = seed
	var end int
	fmt.Println("Seed", seed)
	for _, mapping := range mappings {
		var mappingFound bool
		var mappingFoundIdx int
		for i, edge := range mapping.Edges {
			fmt.Println(edge.Source, start, edge.Source+edge.Increment)
			if edge.Source <= start && start <= edge.Source+edge.Increment {
				mappingFound = true
				mappingFoundIdx = i
				break
			}
		}
		if mappingFound {
			foundEdge := mapping.Edges[mappingFoundIdx]
			offsetFromStart := start - foundEdge.Source
			end = foundEdge.Destination + offsetFromStart
		} else {
			end = start
		}
		fmt.Println(start, " -> ", end)
		start = end
	}
	return end
}

func main() {
	mappings, seeds, err := readMappingsFromFile("./day-5-1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Seeds:", seeds)
	var locations []int
	var minLocation int = seeds[0]
	for _, seed := range seeds {
		foundLocation := getLocationFromSeed(mappings, seed)
		locations = append(locations, foundLocation)
		if foundLocation < minLocation {
			minLocation = foundLocation
		}
	}
	for i, mapping := range mappings {
		fmt.Printf("Mapping %d:\n", i)
		for _, edge := range mapping.Edges {
			fmt.Printf("  %v -> %v, increment: %v\n", edge.Source, edge.Destination, edge.Increment)
		}
	}
	fmt.Println(seeds)
	fmt.Println(locations)
	fmt.Println(minLocation)
}
