package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay8(input string) error {
	s1, err := d8FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d8SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d8FirstPart(input string) (string, error) {
	antennas := make(map[string][][2]int)
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				continue
			}
			if _, ok := antennas[string(c)]; !ok {
				antennas[string(c)] = [][2]int{}
			}
			antennas[string(c)] = append(antennas[string(c)], [2]int{i, j})
		}
	}
	nodes := make(map[[2]int]bool)
	for _, v := range antennas {
		l := len(v)
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				x, y := v[i][0], v[i][1]
				a, b := v[j][0], v[j][1]
				dx := x - a
				dy := y - b
				nx, ny := x+dx, y+dy
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					nodes[[2]int{nx, ny}] = true
				}
				nx, ny = a-dx, b-dy
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					nodes[[2]int{nx, ny}] = true
				}
			}
		}
	}
	return strconv.Itoa(len(nodes)), nil
}

func d8SecondPart(input string) (string, error) {
	antennas := make(map[string][][2]int)
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				continue
			}
			if _, ok := antennas[string(c)]; !ok {
				antennas[string(c)] = [][2]int{}
			}
			antennas[string(c)] = append(antennas[string(c)], [2]int{i, j})
		}
	}
	nodes := make(map[[2]int]bool)
	for _, v := range antennas {
		l := len(v)
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				x, y := v[i][0], v[i][1]
				a, b := v[j][0], v[j][1]
				dx := x - a
				dy := y - b
				gcd := utils.Gcd(dx, dy)
				dx /= gcd
				dy /= gcd
				nx, ny := x, y
				for nx >= 0 && nx < m && ny >= 0 && ny < n {
					nodes[[2]int{nx, ny}] = true
					nx += dx
					ny += dy
				}
				nx, ny = x, y
				for nx >= 0 && nx < m && ny >= 0 && ny < n {
					nodes[[2]int{nx, ny}] = true
					nx -= dx
					ny -= dy
				}

			}
		}
	}
	return strconv.Itoa(len(nodes)), nil
}