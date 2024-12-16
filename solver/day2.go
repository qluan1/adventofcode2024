package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay2(input string) error {
	s1, err := d2FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d2SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d2FirstPart(input string) (string, error) {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetNumbers(line, " ")
		if differenceBounded(nums) && (allIncrease(nums) || allDecrease(nums)) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func d2SecondPart(input string) (string, error) {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetNumbers(line, " ")
		isSafe := false
		for i := 0; i < len(nums); i++ {
			tmp := make([]int, 0)
			tmp = append(tmp, nums[0:i]...)
			tmp = append(tmp, nums[i+1:]...)
			if (allIncrease(tmp) || allDecrease(tmp)) && differenceBounded(tmp) {
				isSafe = true
				break
			}
		}
		if isSafe {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func allIncrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func allDecrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}

func differenceBounded(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := utils.AbsDiff(nums[i], nums[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
