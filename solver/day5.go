package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay5(input string) error {
	s1, err := d5FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d5SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d5FirstPart(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	pairs := make([][]int, 0)
	orders := make([][]int, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		pair := utils.GetNumbers(line, "|")
		pairs = append(pairs, pair)
	}
	for _, line := range strings.Split(parts[1], "\n") {
		order := utils.GetNumbers(line, ",")
		orders = append(orders, order)
	}
	preOrder := make(map[int]map[int]bool)
	for _, pair := range pairs {
		first := pair[0]
		second := pair[1]
		if _, ok := preOrder[second]; !ok {
			preOrder[second] = make(map[int]bool)
		}
		preOrder[second][first] = true
	}
	res := 0
	for _, order := range orders {
		canNotTake := make(map[int]bool)
		for i, num := range order {
			if _, ok := canNotTake[num]; ok {
				break
			}
			for key := range preOrder[num] {
				canNotTake[key] = true
			}
			if i == len(order)-1 {
				res += order[len(order)/2]
			}
		}
	}
	return strconv.Itoa(res), nil
}

func d5SecondPart(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	pairs := make([][]int, 0)
	orders := make([][]int, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		pair := utils.GetNumbers(line, "|")
		pairs = append(pairs, pair)
	}
	for _, line := range strings.Split(parts[1], "\n") {
		order := utils.GetNumbers(line, ",")
		orders = append(orders, order)
	}
	preOrder := make(map[int]map[int]bool)
	for _, pair := range pairs {
		first := pair[0]
		second := pair[1]
		if _, ok := preOrder[second]; !ok {
			preOrder[second] = make(map[int]bool)
		}
		preOrder[second][first] = true
	}
	// get all numbers
	allNumbers := make(map[int]bool)
	for _, order := range orders {
		for _, num := range order {
			allNumbers[num] = true
		}
	}
	for _, pair := range pairs {
		allNumbers[pair[0]] = true
		allNumbers[pair[1]] = true
	}

	res := 0
	for _, order := range orders {
		canNotTake := make(map[int]bool)
		check := 0
		for _, num := range order {
			if _, ok := canNotTake[num]; ok {
				break
			}
			check++
			for key := range preOrder[num] {
				canNotTake[key] = true
			}
		}
		if check == len(order) {
			continue
		}
		i := 0
		for i < len(order) {
			swap := false
			for j := i + 1; j < len(order); j++ {
				if _, ok := preOrder[order[i]][order[j]]; ok {
					order[i], order[j] = order[j], order[i]
					swap = true
					break
				}
			}
			if !swap {
				i++
			}
		}
		res += order[len(order)/2]
	}
	return strconv.Itoa(res), nil
}
