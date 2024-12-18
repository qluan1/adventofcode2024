package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay6(input string) error {
	s1, err := d6FirstPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	s2, err := d6SecondPart(input)
	if err != nil {
		return err
	}
	fmt.Println("Second part:", s2)
	return nil
}

var dirs = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func d6FirstPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	mat := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			mat[i][j] = string(c)
		}
	}
	repeat := make(map[[3]int]bool)
	// pos := [3]int{dir, x, y}
	pos := [3]int{0, 0, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == "#" || mat[i][j] == "." {
				continue
			}
			pos[1] = i
			pos[2] = j
			if mat[i][j] == "^" {
				pos[0] = 0
			} else if mat[i][j] == ">" {
				pos[0] = 1
			} else if mat[i][j] == "v" {
				pos[0] = 2
			} else if mat[i][j] == "<" {
				pos[0] = 3
			}
		}
	}
	for !repeat[pos] {
		repeat[pos] = true
		d := pos[0]
		x := pos[1]
		y := pos[2]
		nx := x + dirs[d][0]
		ny := y + dirs[d][1]
		if nx < 0 || ny < 0 || nx >= m || ny >= n {
			// out of bound
			break
		}
		if mat[nx][ny] == "#" {
			// turn right
			pos[0] = (d + 1) % 4
		} else {
			// move forward
			pos[1] = nx
			pos[2] = ny
		}
	}
	counter := make(map[[2]int]bool)
	for k := range repeat {
		counter[[2]int{k[1], k[2]}] = true
	}

	return strconv.Itoa(len(counter)), nil
}

func d6SecondPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	mat := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			mat[i][j] = string(c)
		}
	}
	// pos := [3]int{dir, x, y}
	pos := [3]int{0, 0, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == "#" || mat[i][j] == "." {
				continue
			}
			pos[1] = i
			pos[2] = j
			if mat[i][j] == "^" {
				pos[0] = 0
			} else if mat[i][j] == ">" {
				pos[0] = 1
			} else if mat[i][j] == "v" {
				pos[0] = 2
			} else if mat[i][j] == "<" {
				pos[0] = 3
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != "." {
				continue
			}
			restore := mat[i][j]
			mat[i][j] = "#"
			cur := [3]int{pos[0], pos[1], pos[2]}
			repeat := make(map[[3]int]bool)
			for {
				repeat[cur] = true
				d := cur[0]
				x := cur[1]
				y := cur[2]
				nx := x + dirs[d][0]
				ny := y + dirs[d][1]
				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					// out of bound
					break
				}
				if mat[nx][ny] == "#" {
					// turn right
					cur[0] = (d + 1) % 4
				} else {
					// move forward
					cur[1] = nx
					cur[2] = ny
				}
				if repeat[cur] {
					res++
					break
				}
			}
			mat[i][j] = restore
		}
	}
	return strconv.Itoa(res), nil
}