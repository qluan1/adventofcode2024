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
		fmt.Println("Loging is enabled")
	} else {
		fmt.Println("Loging is disabled")
	}
	fmt.Printf("Using solver for day %d, input is provided from %s\n", day, file)

	solverFunc := solvers[day]
	err := solverFunc(input)
	if err != nil {
		fmt.Println(err)
	}
}
