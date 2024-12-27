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

func d14SecondPart(input string) (string, error) {
	return "", nil
}