package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/qluan1/adventofcode2024/solver"
	"github.com/qluan1/adventofcode2024/utils"
)

var day uint
var file string
var input string

var solvers = map[uint]func(string) error{
	1: solver.SolveDay1,
	2: solver.SolveDay2,
	3: solver.SolveDay3,
	4: solver.SolveDay4,
	5: solver.SolveDay5,
	6: solver.SolveDay6,
	7: solver.SolveDay7,
	8: solver.SolveDay8,
	9: solver.SolveDay9,
	10: solver.SolveDay10,
	11: solver.SolveDay11,
	12: solver.SolveDay12,
	13: solver.SolveDay13,
}

func init() {
	flag.BoolVar(&utils.EnableLogging, "log", false, "Enable Logging")
	flag.UintVar(&day, "day", 1, "Solver for day")
	flag.StringVar(&file, "file", "", "Input file")
	flag.Parse()
	fmt.Println(file)
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input = string(data)
}
func main() {
	if _, ok := solvers[day]; !ok {
		fmt.Printf("Solver for day %d not found\n", day)
	}
	if utils.EnableLogging {
		fmt.Println("Logging is enabled")
	} else {
		fmt.Println("Logging is disabled")
	}
	fmt.Printf("Using solver for day %d, input is provided from %s\n", day, file)

	solverFunc := solvers[day]
	err := solverFunc(input)
	if err != nil {
		fmt.Println(err)
	}
}
