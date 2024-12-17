package solver

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func d1FirstPart(input string) (string, error) {
	column1 := []int{}
	column2 := []int{}
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetNumbers(line, " ")
		column1 = append(column1, nums[0])
		column2 = append(column2, nums[1])
	}
	sort.Ints(column1)
	sort.Ints(column2)
	res := 0	
	for i := 0; i < len(column1); i++ {
		res += utils.AbsDiff(column1[i], column2[i])
	}

	return strconv.Itoa(res), nil
}

func d1SecondPart(input string) (string, error) {
	column1 := []int{}
	count := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetNumbers(line, " ")
		column1 = append(column1, nums[0])
		count[nums[1]]++
	}
	total := 0
	for _, v := range column1 {
		total += v * count[v]
	}
	return strconv.Itoa(total), nil	
}

func SolveDay1(input string) error {
	s1, err := d1FirstPart(input)
	if err != nil {
		return err
	}
	s2, err := d1SecondPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	fmt.Println("Second part:", s2)
	return nil
}
