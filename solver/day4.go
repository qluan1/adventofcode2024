package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay4(input string) error {
	s1, err := d4FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d4SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d4FirstPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	mat := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			mat[i][j] = string(c)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		horizontal := mat[i]
		count += d4Total(horizontal)
		count += d4Total(d4Reverse(horizontal))
	}
	for j := 0; j < n; j++ {
		vertical := make([]string, 0)
		for i := 0; i < m; i++ {
			vertical = append(vertical, mat[i][j])
		}
		count += d4Total(vertical)
		count += d4Total(d4Reverse(vertical))
	}
	// diagonal
	for i := 0; i < m; i++ {
		diagonal := make([]string, 0)
		for k := 0; i+k < m && k < n; k++ {
			diagonal = append(diagonal, mat[i+k][k])
		}
		count += d4Total(diagonal)
		count += d4Total(d4Reverse(diagonal))
	}
	for j := 1; j < n; j++ {
		diagonal := make([]string, 0)
		for k := 0; j+k < n && k < m; k++ {
			diagonal = append(diagonal, mat[k][j+k])
		}
		count += d4Total(diagonal)
		count += d4Total(d4Reverse(diagonal))
	}
	// anti-diagonal
	for i := 0; i < m; i++ {
		antiDiagonal := make([]string, 0)
		for k := 0; i-k >= 0 && k < n; k++ {
			antiDiagonal = append(antiDiagonal, mat[i-k][k])
		}
		count += d4Total(antiDiagonal)
		count += d4Total(d4Reverse(antiDiagonal))
	}
	for j := 1; j < n; j++ {
		antiDiagonal := make([]string, 0)
		for k := 0; j+k < n && m - 1 - k > -1; k++ {
			antiDiagonal = append(antiDiagonal, mat[m-1-k][j+k])
		}
		count += d4Total(antiDiagonal)
		count += d4Total(d4Reverse(antiDiagonal))
	}
	return strconv.Itoa(count), nil
}

func d4Total(s1 []string) int {
	count := 0
	for i := range s1 {
		if i + 3 >= len(s1) {
			break
		}
		if s1[i] == "X" && s1[i+1] == "M" && s1[i+2] == "A" && s1[i+3] == "S" {
			count++
		}
	}
	return count
}

func d4Reverse(s []string) []string {
	ss := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ss[i] = s[len(s)-1-i]
	}
	return ss
}

func d4SecondPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	mat := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			mat[i][j] = string(c)
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count += d4CountSquare(mat, i, j)
		}
	}
	return strconv.Itoa(count), nil
}

func d4CountSquare(mat [][]string, i, j int) int {
	count := 0
	m := len(mat)
	n := len(mat[0])
	if i - 1 < 0 || i + 1 >= m || j - 1 < 0 || j + 1 >= n {
		return 0
	}
	if mat[i][j] != "A" {
		return 0
	}
	quad := [][4]int{
		{i-1, j-1, i+1, j+1},
		{i+1, j-1, i-1, j+1},
	}

	for _, q := range quad {
		x, y, xx, yy := q[0], q[1], q[2], q[3]
		if mat[x][y] == "M" && mat[xx][yy] == "S" {
			count++
		}
		if mat[x][y] == "S" && mat[xx][yy] == "M" {
			count++
		}
	}
	if count >= 2 {
		return 1
	}
	return 0
}