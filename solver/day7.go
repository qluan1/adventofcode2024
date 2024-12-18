package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay7(input string) error {
	s1, err := d7FirstPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	s2, err := d7SecondPart(input)
	if err != nil {
		return err
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d7FirstPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	test := make([]int, 0)
	eqn := make([][]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		v, _ := strconv.Atoi(parts[0])
		test = append(test, v)
		nums := utils.GetNumbers(parts[1], " ")
		eqn = append(eqn, nums)
	}
	res := 0
	for i := 0; i < m; i++ {
		if d7Helper(eqn[i], test[i]) {
			res += test[i]
		}
	}
	return strconv.Itoa(res), nil
}

func d7SecondPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	test := make([]int, 0)
	eqn := make([][]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		v, _ := strconv.Atoi(parts[0])
		test = append(test, v)
		nums := utils.GetNumbers(parts[1], " ")
		eqn = append(eqn, nums)
	}
	res := 0
	for i := 0; i < m; i++ {
		if d7Helper2(eqn[i], test[i]) {
			res += test[i]
		}
	}
	return strconv.Itoa(res), nil
}

func d7Helper(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	if l == 1 {
		return nums[0] == target
	}
	// assume not overflowing
	size := 1
	for i := 0; i < l-1; i++ {
		size *= 2
	}
	for i := 0; i < size; i++ {
		d := i
		v := nums[0]
		for j := 1; j < l; j++ {
			if d%2 == 0 {
				v += nums[j]
			} else {
				v *= nums[j]
			}
			d /= 2
		}
		if v == target {
			return true
		}
	}
	return false
}

func d7Helper2(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	if l == 1 {
		return nums[0] == target
	}
	// assume not overflowing
	size := 1
	for i := 0; i < l-1; i++ {
		size *= 3
	}
	for i := 0; i < size; i++ {
		d := i
		v := nums[0]
		for j := 1; j < l; j++ {
			if d%3 == 0 {
				v += nums[j]
			} else if d%3 == 1 {
				v *= nums[j]
			} else {
				s1 := strconv.Itoa(v)
				s2 := strconv.Itoa(nums[j])
				v, _ = strconv.Atoi(s1 + s2)
			}
			d /= 3
		}
		if v == target {
			return true
		}
	}
	return false

}