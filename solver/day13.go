package solver

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay13(input string) error {
	s1, err := d13FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d13SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d13ParseBlock(input string) [6]int {
	lines := strings.Split(input, "\n")
	var a, b, c, d, e, f int
	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &a, &b)
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &c, &d)
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &e, &f)
	return [6]int{a, b, c, d, e, f}
}

func d13Solve(nums [6]int) (bool, int) {
	x1, y1, x2, y2, x, y := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]
	res := math.MaxInt64
	iter := math.MaxInt64
	iter = utils.Min(iter, x/x1)
	iter = utils.Min(iter, y/y1)
	for i := 0; i <= iter; i++ {
		if (x-i*x1)%x2 == 0 && (y-i*y1)%y2 == 0  && (x-i*x1)/x2 == (y-i*y1)/y2 {
			res = utils.Min(res, 3*i + (x-i*x1)/x2)
		}
	}
	if res != math.MaxInt64 {
		return true, res
	}
	return false, 0
}

func d13Solve2(nums [6]int) (bool, int) {
	x1, y1, x2, y2, x, y := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]
	x, y = x + 10000000000000, y + 10000000000000
	if x1 * y2 - x2 * y1 == 0 { // multiple solution
		lcm := x1 * x2 / utils.Gcd(x1, x2)
		if 3 * x2 >= x1 {
			for i := 0; i <= lcm/x1; i++ {
				if (x - i*x1) % x2 == 0 && (y - i*y1) % y2 == 0 && (x - i*x1) / x2 == (y - i*y1) / y2 {
					return true, 3 * i + (x - i*x1) / x2
				}
			}
		} else {
			for i := 0; i <= lcm/x2; i++ {
				if (x - i*x2) % x1 == 0 && (y - i*y2) % y1 == 0 && (x - i*x2) / x1 == (y - i*y2) / y1 {
					return true, i
				}
			}
		}
		return false, 0
	}
	det := x1 * y2 - x2 * y1
	ix1, ix2, iy1, iy2 := y2, -x2, -y1, x1
	// check int condition
	if (ix1*x + ix2*y) % det != 0 || (iy1*x + iy2*y) % det != 0 {
		return false, 0
	}
	r1 := (ix1*x + ix2*y) / det
	r2 := (iy1*x + iy2*y) / det
	if r1 < 0 || r2 < 0 {
		return false, 0
	}
	return true, 3*r1 + r2
}

func d13FirstPart(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	res := 0
	for _, block := range blocks {
		b := d13ParseBlock(block)
		ok, v := d13Solve(b)
		if ok {
			res += v
		}
	}
	return strconv.Itoa(res), nil
}

func d13SecondPart(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	res := 0
	for _, block := range blocks {
		b := d13ParseBlock(block)
		ok, v := d13Solve2(b)
		if ok {
			res += v
		}
	}
	return strconv.Itoa(res), nil
}
