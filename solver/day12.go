package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

func SolveDay12(input string) error {
	s1, err := d12FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d12SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d12ParseInput(input string) ([][]string, error) {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	mat := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			mat[i][j] = string(c)
		}
	}
	return mat, nil
}

func d12FirstPart(input string) (string, error) {
	mat, _ := d12ParseInput(input)
	m, n := len(mat), len(mat[0])
	used := map[[2]int]bool{}
	regions := [][][2]int{}	
	var dfs func(int, int, *map[[2]int]bool)
	dfs = func(i, j int, visited *map[[2]int]bool) {
		for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			ni, nj := i+d[0], j+d[1]
			if ni < 0 || nj < 0 || ni >= len(mat) || nj >= len(mat[0]) {
				continue
			}
			if mat[ni][nj] == mat[i][j] && !(*visited)[[2]int{ni, nj}] {
				(*visited)[[2]int{ni, nj}] = true
				dfs(ni, nj, visited)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if used[[2]int{i, j}] {
				continue
			}
			visited := map[[2]int]bool{}
			visited[[2]int{i, j}] = true
			dfs(i, j, &visited)
			region := [][2]int{}
			for k := range visited {
				region = append(region, k)
				used[k] = true
			}
			regions = append(regions, region)
		}
	}
	res := 0
	for _, region := range regions {
		d := make(map[[2]int]bool)
		perimeter := 0
		for _, p := range region {
			d[p] = true
		}
		for _, p := range region {
			for _, dd := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				np := [2]int{p[0] + dd[0], p[1] + dd[1]}
				if _, ok := d[np]; !ok {
					perimeter++
				}
			}
		}
		res += perimeter * len(region)
	}

	return strconv.Itoa(res), nil
}

func d12SecondPart(input string) (string, error) {
	mat, _ := d12ParseInput(input)
	m, n := len(mat), len(mat[0])
	used := map[[2]int]bool{}
	regions := [][][2]int{}	
	var dfs func(int, int, *map[[2]int]bool)
	dfs = func(i, j int, visited *map[[2]int]bool) {
		for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			ni, nj := i+d[0], j+d[1]
			if ni < 0 || nj < 0 || ni >= len(mat) || nj >= len(mat[0]) {
				continue
			}
			if mat[ni][nj] == mat[i][j] && !(*visited)[[2]int{ni, nj}] {
				(*visited)[[2]int{ni, nj}] = true
				dfs(ni, nj, visited)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if used[[2]int{i, j}] {
				continue
			}
			visited := map[[2]int]bool{}
			visited[[2]int{i, j}] = true
			dfs(i, j, &visited)
			region := [][2]int{}
			for k := range visited {
				region = append(region, k)
				used[k] = true
			}
			regions = append(regions, region)
		}
	}
	res := 0
	for _, region := range regions {
		d := make(map[[2]int]bool)
		for _, p := range region {
			d[p] = true
		}
		hBoundaries := map[[3]int]bool{} // count (x, y, dir) top is bound
		vBoundaries := map[[3]int]bool{} // count (x, y, dir) left is bound
		for _, p := range region {
			x, y := p[0], p[1]
			// horizontal
			for _, dd := range [][2]int{{1, 0}, {-1, 0}} {
				nx := x + dd[0]
				if _, ok := d[[2]int{nx, y}]; !ok {
					hBoundaries[[3]int{utils.Max(x, nx), y, dd[0]}] = true
				}
			}
			// vertical
			for _, dd := range [][2]int{{0, 1}, {0, -1}} {
				ny := y + dd[1]
				if _, ok := d[[2]int{x, ny}]; !ok {
					vBoundaries[[3]int{x, utils.Max(y, ny), dd[1]}] = true
				}
			}
		}
		sides := 0
		// count the number of sides by depleting the boundaries
		usedHBound := map[[3]int]bool{}
		for k := range hBoundaries {
			if usedHBound[k] {
				continue
			}
			x, y, dir := k[0], k[1], k[2]
			// seek left till end
			left := y
			for hBoundaries[[3]int{x, left, dir}] && !usedHBound[[3]int{x, left, dir}] {
				usedHBound[[3]int{x, left, dir}] = true
				left--
			}
			// seek right till end
			right := y+1
			for hBoundaries[[3]int{x, right, dir}] && !usedHBound[[3]int{x, right, dir}] {
				usedHBound[[3]int{x, right, dir}] = true
				right++
			}
			sides++
		}
		usedVBound := map[[3]int]bool{}
		for k := range vBoundaries {
			if usedVBound[k] {
				continue
			}
			x, y, dir := k[0], k[1], k[2]
			// seek up till end
			up := x
			for vBoundaries[[3]int{up, y, dir}] && !usedVBound[[3]int{up, y, dir}] {
				usedVBound[[3]int{up, y, dir}] = true
				up--
			}
			// seek down till end
			down := x+1
			for vBoundaries[[3]int{down, y, dir}] && !usedVBound[[3]int{down, y, dir}] {
				usedVBound[[3]int{down, y, dir}] = true
				down++
			}
			sides++
		}
		res += sides * len(region)
	}

	return strconv.Itoa(res), nil
}