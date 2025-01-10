package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay17(input string) error {
	s1, err := d17FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d17SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d17FirstPart(input string) (string, error) {
	table := map[int]func(*[3]int, *[]int, *int, *[]int){
		0: d17adv,
		1: d17bxl,
		2: d17bst,
		3: d17jnz,
		4: d17bxc,
		5: d17out,
		6: d17bdv,
		7: d17cdv,
	}
	registries, instructions, err := day17ParseInput(input)
	if err != nil {
		return "", err
	}
	pointer := 0
	output := make([]int, 0)
	for pointer < len(instructions) {
		table[instructions[pointer]](&registries, &instructions, &pointer, &output)
	}
	strs := make([]string, len(output))
	for i, v := range output {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ","), nil
}

func d17SecondPart(input string) (string, error) {
	return "", nil
}

func day17ParseInput(input string) ([3]int, []int, error) {
	lines := strings.Split(input, "\n")	

	registries := [3]int{}
	for i := 0; i < 3; i++ {
		a, err := strconv.Atoi(strings.Split(lines[i], ": ")[1])
		if err != nil {
			return registries, nil, err
		}
		registries[i] = a
	}

	nums := utils.GetNumbers(strings.Split(lines[4], ": ")[1], ",")
	return registries, nums, nil
}

func d17combo(registries *[3]int, operand int) int {
	if operand <=3 {
		return operand
	}
	if operand <= 6 {
		return registries[operand-4]
	}
	return -1 // will not occur
}

func d17adv(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	comboOperand := d17combo(registries, operand)
	numerator := registries[0]
	denominator := utils.Pow(2, comboOperand)	
	(*registries)[0] = numerator / denominator
	(*pointer)++
	(*pointer)++
}

func d17bxl(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	(*registries)[1] = registries[1] ^ operand
	(*pointer)++
	(*pointer)++
}

func d17bst(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	(*registries)[1] = d17combo(registries, operand) % 8
	(*pointer)++
	(*pointer)++
}

func d17jnz(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	if registries[0] == 0 {
		(*pointer)++
		(*pointer)++
		return
	}
	(*pointer) = operand
}

func d17bxc(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	(*registries)[1] = registries[1] ^ registries[2]
	(*pointer)++
	(*pointer)++
	return
}

func d17out(registries *[3]int, instructions *[]int, pointer *int, output *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	comboOperand := d17combo(registries, operand)
	*output = append(*output, comboOperand%8)
	(*pointer)++
	(*pointer)++
}

func d17bdv(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	comboOperand := d17combo(registries, operand)
	numerator := registries[0]
	denominator := utils.Pow(2, comboOperand)	
	(*registries)[1] = numerator / denominator
	(*pointer)++
	(*pointer)++
}

func d17cdv(registries *[3]int, instructions *[]int, pointer *int, _ *[]int) {
	if *pointer + 1 >= len(*instructions) {
		(*pointer)++
		return
	}
	operand := (*instructions)[*pointer + 1]
	comboOperand := d17combo(registries, operand)
	numerator := registries[0]
	denominator := utils.Pow(2, comboOperand)	
	(*registries)[2] = numerator / denominator
	(*pointer)++
	(*pointer)++
}