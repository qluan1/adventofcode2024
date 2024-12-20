package solver

import (
	"fmt"
	"strconv"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay11(input string) error {
	s1, err := d11FirstPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	s2, err := d11SecondPart(input)
	if err != nil {
		return err
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d11FirstPart(input string) (string, error) {
	nums := utils.GetNumbers(input, " ")
	var nxt []int
	for i := 0; i < 25; i++ {
		nxt = []int{}
		for _, num := range nums {
			option := d11GetOption(num)
			if option == 0 {
				nxt = append(nxt, 1)
				continue
			}
			if option == 1 {
				v1, v2 := d11BreakNumber(num)
				nxt = append(nxt, v1)
				nxt = append(nxt, v2)
				continue
			}
			if option == 2 {
				nxt = append(nxt, num * 2024)
			}
		}
		nums = nxt
	}
	return strconv.Itoa(len(nxt)), nil
}

func d11SecondPart(input string) (string, error) {
	nums := utils.GetNumbers(input, " ")
	cached := make(map[int]map[int]int)
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	for i := 0; i < 3; i++ {
		newCounter := make(map[int]int)
		for num, count := range counter {
			if _, ok := cached[num]; !ok {
				cached[num] = d11Get25Blinks([]int{num})
			}
			for k, v := range cached[num] {
				newCounter[k] += v * count
			}
		}
		counter = newCounter
	}
	res := 0
	for _, v := range counter {
		res += v
	}
	return strconv.Itoa(res), nil
}

func d11GetOption(num int) int {
	if num == 0 {
		return 0
	}
	if len(strconv.Itoa(num)) % 2 == 0 {
		return 1
	}
	return 2
}

func d11BreakNumber(num int) (int, int) {
	arr := []int{}
	for num > 0 {
		arr = append(arr, num % 10)
		num /= 10
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
	}
	v1, v2 := 0, 0
	base := 1
	for i := 0; i < len(arr)/2; i++ {
		v1 += arr[len(arr)/2 - 1 -i] * base
		v2 += arr[len(arr) - 1 - i] * base
		base *= 10
	}
	return v1, v2
}

func d11Get25Blinks(nums []int) map[int]int {
	var nxt []int
	for i := 0; i < 25; i++ {
		nxt = []int{}
		for _, num := range nums {
			option := d11GetOption(num)
			if option == 0 {
				nxt = append(nxt, 1)
				continue
			}
			if option == 1 {
				v1, v2 := d11BreakNumber(num)
				nxt = append(nxt, v1)
				nxt = append(nxt, v2)
				continue
			}
			if option == 2 {
				nxt = append(nxt, num * 2024)
			}
		}
		nums = nxt
	}
	res := make(map[int]int)
	for _, num := range nxt {
		res[num]++
	}
	return res
}
