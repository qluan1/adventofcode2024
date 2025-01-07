package solver

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveDay14(input string) error {
	s1, err := d14FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d14SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func readRobot(input string) (int, int, int, int) {
	var x, y, vx, vy int
	fmt.Sscanf(input, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
	return x, y, vx, vy
}

func getNextPos(x, y, vx, vy, xMax, yMax int) (int, int) {
	return (x + vx + xMax) % xMax, (y + vy + yMax) % yMax
}

func d14FirstPart(input string) (string, error) {
	robots := make([][4]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		x, y, vx, vy := readRobot(line)
		robots = append(robots, [4]int{x, y, vx, vy})
	}
	wide, tall := 101, 103
	for i := 0; i < 100; i++ {
		for i, robot := range robots {
			x, y := getNextPos(robot[0], robot[1], robot[2], robot[3], wide, tall)
			robots[i] = [4]int{x, y, robot[2], robot[3]}
		}
	}
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		x, y := robot[0], robot[1]
		if x < wide/2 && y < tall/2 {
			q1++
		} else if x > wide/2 && y < tall/2 {
			q2++
		} else if x < wide/2 && y > tall/2 {
			q3++
		} else if x > wide/2 && y > tall/2{
			q4++
		}
	}
	return strconv.Itoa(q1*q2*q3*q4), nil
}

func d14CountConnectedComponents(robots [][4]int) int {
	d := make(map[[2]int]int)	
	for i, robot := range robots {
		d[[2]int{robot[0], robot[1]}] = i
	}
	arr := make([]int, len(robots))
	for i := range arr {
		arr[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}
	union := func(x, y int) {
		arr[find(x)] = find(y)
	}
	for i, robot := range robots {
		x, y := robot[0], robot[1]
		if _, ok := d[[2]int{x+1, y}]; ok {
			union(i, d[[2]int{x+1, y}])
		}
		if _, ok := d[[2]int{x, y+1}]; ok {
			union(i, d[[2]int{x, y+1}])
		}
	}
	for i := range arr {
		find(i)
	}
	components := make(map[int]bool)
	for _, v := range arr {
		components[v] = true
	}
	return len(components)
}

func d14SecondPart(input string) (string, error) {
	robots := make([][4]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		x, y, vx, vy := readRobot(line)
		robots = append(robots, [4]int{x, y, vx, vy})
	}
	wide, tall := 101, 103
	step := -1
	for i := 0; i < 100000; i++ {
		connectedComponents := d14CountConnectedComponents(robots)
		if connectedComponents < len(robots)/2 {
			step = i	
			break
		}
		for i, robot := range robots {
			x, y := getNextPos(robot[0], robot[1], robot[2], robot[3], wide, tall)
			robots[i] = [4]int{x, y, robot[2], robot[3]}
		}
	}
	if step == -1 {
		return "", fmt.Errorf("no solution found by the 100000th step")
	}
	return strconv.Itoa(step), nil
}