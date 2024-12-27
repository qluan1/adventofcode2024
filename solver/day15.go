package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay15(input string) error {
	s1, err := d15FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d15SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d15ParseInput(input string) ([][]int, [][2]int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	arr := strings.Split(input, "\n\n")
	mp := strings.Split(arr[0], "\n")
	m, n := len(mp), len(mp[0])	
	mat := utils.GetIntMat(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == '#' {
				mat[i][j] = -1
			} else if mp[i][j] == '.' {
				mat[i][j] = 0
			} else if mp[i][j] == 'O' {
				mat[i][j] = 1
			} else {
				mat[i][j] = 2 // robot
			}
		}
	}
	instructions := make([][2]int, 0)
	for _, line := range strings.Split(arr[1], "\n") {
		for _, chr := range line {
			if chr == '^' {
				instructions = append(instructions, [2]int{-1, 0})
			} else if chr == 'v' {
				instructions = append(instructions, [2]int{1, 0})
			} else if chr == '<' {
				instructions = append(instructions, [2]int{0, -1})
			} else {
				instructions = append(instructions, [2]int{0, 1})
			}
		}
	}
	return mat, instructions
}

func d15Move(mat *[][]int, x, y, dx, dy int) (int, int) {
	curX, curY := x, y
	for (*mat)[curX][curY] == 2 || (*mat)[curX][curY] == 1 {
		curX += dx
		curY += dy
	}
	if (*mat)[curX][curY] == 0 {
		(*mat)[curX][curY] = 1
		(*mat)[x][y] = 0
		(*mat)[x+dx][y+dy] = 2
		return x+dx, y+dy
	}
	return x, y
}

func d15FirstPart(input string) (string, error) {
	mat, instructions := d15ParseInput(input)
	m, n := len(mat), len(mat[0])
	x, y := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 2 {
				x, y = i, j
				break
			}
		}
	}
	for _, instruction := range instructions {
		x, y = d15Move(&mat, x, y, instruction[0], instruction[1])
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				res += 100 * i + j
			}
		}
	}
	return strconv.Itoa(res), nil
}

func d15BuildMat(mat [][]int) [][]string {
	m, n := len(mat), len(mat[0])
	pm := utils.GetStrMat(m, 2*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == -1 {
				pm[i][2*j] = "#"
				pm[i][2*j+1] = "#"
			} else if mat[i][j] == 0 {
				pm[i][2*j] = "."
				pm[i][2*j+1] = "."
			} else if mat[i][j] == 1 {
				pm[i][2*j] = "["
				pm[i][2*j+1] = "]"
			} else {
				pm[i][2*j] = "@"
				pm[i][2*j+1] = "."
			}
		}
	}
	return pm
}

func d15SecondPart(input string) (string, error) {
	mat, instructions := d15ParseInput(input)
	sm := d15BuildMat(mat)
	m, n := len(sm), len(sm[0])
	x, y := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if sm[i][j] == "@" {
				x, y = i, j
				break
			}
		}
	}
	for _, instruction := range instructions {
		dx, dy := instruction[0], instruction[1]
		if d15CanPush(sm, x, y, dx, dy) {
			d15Push(&sm, x, y, dx, dy)
			x, y = x+dx, y+dy
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if sm[i][j] == "[" {
				res += 100 * i + j
			}
		}
	}
	return strconv.Itoa(res), nil
}

func d15CanPush(mat [][]string, x, y, dx, dy int) bool {
	if mat[x+dx][y+dy] == "#" {
		return false
	}
	if mat[x+dx][y+dy] == "." {
		return true
	}
	if dy == 0 {
		if mat[x+dx][y+dy] == "[" {
			return d15CanPush(mat, x+dx, y+dy+1, dx, dy) && d15CanPush(mat, x+dx, y+dy, dx, dy)
		}
		if mat[x+dx][y+dy] == "]" {
			return d15CanPush(mat, x+dx, y+dy-1, dx, dy) && d15CanPush(mat, x+dx, y+dy, dx, dy)
		}
	}
	if dx == 0 {
		return d15CanPush(mat, x+dx, y+dy, dx, dy)
	}
	return false
}

func d15Push(mat *[][]string, x, y, dx, dy int) {
	if (*mat)[x+dx][y+dy] == "." {
		(*mat)[x+dx][y+dy] = (*mat)[x][y]
		(*mat)[x][y] = "."
		return
	}
	if dy == 0 {
		if (*mat)[x+dx][y+dy] == "[" {
			d15Push(mat, x+dx, y+dy, dx, dy)
			d15Push(mat, x+dx, y+dy+1, dx, dy)
			(*mat)[x+dx][y+dy] =(*mat)[x][y]
			(*mat)[x][y] = "."
			return
		}
		if (*mat)[x+dx][y+dy] == "]" {
			d15Push(mat, x+dx, y+dy, dx, dy)
			d15Push(mat, x+dx, y+dy-1, dx, dy)
			(*mat)[x+dx][y+dy] =(*mat)[x][y]
			(*mat)[x][y] = "."
			return
		}
	}
	d15Push(mat, x+dx, y+dy, dx, dy)
	(*mat)[x+dx][y+dy] = (*mat)[x][y]
	(*mat)[x][y] = "."
}
