package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay10(input string) error {
	s1, err := d10FirstPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	s2, err := d10SecondPart(input)
	if err != nil {
		return err
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d10ParseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	mat := utils.GetIntMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			v, _ := strconv.Atoi(string(c))
			mat[i][j] = v
		}
	}
	return mat, nil
}

func d10FirstPart(input string) (string, error) {
	mat, _ := d10ParseInput(input)
	m, n := len(mat), len(mat[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				continue
			}
			target := 1
			cur := [][2]int{{i, j}}
			for len(cur) > 0 {
				next := [][2]int{}
				used := make(map[[2]int]bool)
				for _, p := range cur {
					x, y := p[0], p[1]
					for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
						dx, dy := dir[0], dir[1]
						nx, ny := x+dx, y+dy
						if nx >= 0 && nx < m &&
							ny >= 0 && ny < n &&
							mat[nx][ny] == target &&
							!used[[2]int{nx, ny}] {
							next = append(next, [2]int{nx, ny})
							mat[nx][ny] = target
							used[[2]int{nx, ny}] = true
						}
					}
				}
				if target == 9 {
					res += len(next)
					break
				}
				target++
				cur = next
			}
		}
	}
	return strconv.Itoa(res), nil
}

func d10SecondPart(input string) (string, error) {
	mat, _ := d10ParseInput(input)
	m, n := len(mat), len(mat[0])
	res := 0
	dp := utils.GetIntMat(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		if mat[i][j] == 9 {
			dp[i][j] = 1
			return 1
		}
		count := 0	
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			dx, dy := dir[0], dir[1]
			nx, ny := i+dx, j+dy
			if nx >= 0 && nx < m &&
				ny >= 0 && ny < n &&
				mat[nx][ny] == mat[i][j] + 1 {
					count += dfs(nx, ny)
			}
		}
		dp[i][j] = count
		return count
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				res += dp[i][j]
			}
		}
	}
	return strconv.Itoa(res), nil
}