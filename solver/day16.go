package solver

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/qluan1/adventofcode2024/utils"
)

var d16Dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Grid struct {
	val [3]int
	key int
}

func (g *Grid) IsLess(x interface{}) bool {
	otherGrid := x.(*Grid)
	return g.key < otherGrid.key
}

func SolveDay16(input string) error {
	s1, err := d16FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d16SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d16FirstPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	maze := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			maze[i][j] = string(c)
		}
	}
	startX, startY := 0, 0
	endX, endY := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maze[i][j] == "S" {
				startX, startY = i, j
			}
			if maze[i][j] == "E" {
				endX, endY = i, j
			}
		}
	}
	// dijkstra
	dist := make(map[[3]int]int)
	pq := &utils.MyHeap{}
	heap.Init(pq)
	heap.Push(pq, &Grid{val: [3]int{startX, startY, 0}, key: 0})
	
	reachedEnd := func () (int, bool) {
		if len(*pq) == 0 {
			return -1, true
		}
		for i := 0; i < 4; i++ {
			if v, ok := dist[[3]int{endX, endY, i}]; ok {
				return v, true
			}
		}
		return -1, false
	}
	for _, ok := reachedEnd(); !ok; _, ok = reachedEnd() {
		grid := heap.Pop(pq).(*Grid)
		if _, ok := dist[grid.val]; ok {
			continue
		}
		dist[grid.val] = grid.key
		x, y, dir := grid.val[0], grid.val[1], grid.val[2]
		// move forward
		nx, ny := x + d16Dirs[dir][0], y + d16Dirs[dir][1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && maze[nx][ny] != "#" {
			heap.Push(pq, &Grid{val: [3]int{nx, ny, dir}, key: grid.key + 1})
		}
		// turn one way
		heap.Push(pq, &Grid{val: [3]int{x, y, (dir + 3) % 4}, key: grid.key + 1000})
		// turn the other way
		heap.Push(pq, &Grid{val: [3]int{x, y, (dir + 1) % 4}, key: grid.key + 1000})
	}
	v, _ := reachedEnd()
	return strconv.Itoa(v), nil
}

func d16SecondPart(input string) (string, error) {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	maze := utils.GetStrMat(m, n)
	for i, line := range lines {
		for j, c := range line {
			maze[i][j] = string(c)
		}
	}
	startX, startY := 0, 0
	endX, endY := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maze[i][j] == "S" {
				startX, startY = i, j
			}
			if maze[i][j] == "E" {
				endX, endY = i, j
			}
		}
	}
	// dijkstra
	dist := make(map[[3]int]int)
	pq := &utils.MyHeap{}
	heap.Init(pq)
	heap.Push(pq, &Grid{val: [3]int{startX, startY, 0}, key: 0})
	
	reachedEnd := func () (int, bool) {
		if len(*pq) == 0 {
			return -1, true
		}
		for i := 0; i < 4; i++ {
			if v, ok := dist[[3]int{endX, endY, i}]; ok {
				return v, true
			}
		}
		return -1, false
	}
	for v, ok := reachedEnd(); !ok || pq.Peek().(*Grid).key <= v; v, ok = reachedEnd() {
		grid := heap.Pop(pq).(*Grid)
		if _, ok := dist[grid.val]; ok {
			continue
		}
		dist[grid.val] = grid.key
		x, y, dir := grid.val[0], grid.val[1], grid.val[2]
		// move forward
		nx, ny := x + d16Dirs[dir][0], y + d16Dirs[dir][1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && maze[nx][ny] != "#" {
			heap.Push(pq, &Grid{val: [3]int{nx, ny, dir}, key: grid.key + 1})
		}
		// turn one way
		heap.Push(pq, &Grid{val: [3]int{x, y, (dir + 3) % 4}, key: grid.key + 1000})
		// turn the other way
		heap.Push(pq, &Grid{val: [3]int{x, y, (dir + 1) % 4}, key: grid.key + 1000})
	}
	minCost, _ := reachedEnd()
	used := make(map[[2]int]bool)
	var backtrack func([3]int)
	backtrack = func (pos [3]int) {
		x, y, dir := pos[0], pos[1], pos[2]
		used[[2]int{x, y}] = true
		// move backward
		nx, ny := x - d16Dirs[dir][0], y - d16Dirs[dir][1]
		if v, ok := dist[[3]int{nx, ny, dir}]; ok && v == dist[pos] - 1 {
			backtrack([3]int{nx, ny, dir})
		}
		// turn one way
		if v, ok := dist[[3]int{x, y, (dir + 3) % 4}]; ok && v == dist[pos] - 1000 {
			backtrack([3]int{x, y, (dir + 3) % 4})
		}
		// turn the other way
		if v, ok := dist[[3]int{x, y, (dir + 1) % 4}]; ok && v == dist[pos] - 1000 {
			backtrack([3]int{x, y, (dir + 1) % 4})
		}
	}
	for i := 0; i < 4; i++ {
		if v, ok := dist[[3]int{endX, endY, i}]; ok && v == minCost {
			backtrack([3]int{endX, endY, i})
		}
	}
	return strconv.Itoa(len(used)), nil
}
