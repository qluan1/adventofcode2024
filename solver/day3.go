package solver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolveDay3(input string) error {
	s1, err := d3FirstPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First part:", s1)
	s2, err := d3SecondPart(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d3FirstPart(input string) (string, error) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	arr := re.FindAllString(input, -1)
	sum := 0
	for _, s := range arr {
		v, err := helper(s)
		if err != nil {
			return "", err
		}
		sum += v
	}
	return strconv.Itoa(sum), nil
}

func d3SecondPart(input string) (string, error) {
	sum := 0
	dos := strings.Split(input, "do()")
	for _, do := range dos {
		donts := strings.Split(do, "don't()")
		if len(donts) > 0 {
			v, err := d3FirstPart(donts[0])
			if err != nil {
				return "", err
			}
			vv, err := strconv.Atoi(v)
			if err != nil {
				return "", err
			}
			sum += vv
		}
	}
	return strconv.Itoa(sum), nil
}

func helper(s string) (int, error) {
	r := []rune(s)
	i1 := -1
	i2 := -1
	i3 := -1
	for i, c := range r {
		if c == '(' {
			i1 = i
		}
		if c == ',' {
			i2 = i
		}
		if c == ')' {
			i3 = i
		}
	} 
	if i1 == -1 || i2 == -1 || i3 == -1 {
		return 0, fmt.Errorf("invalid input")
	}
	v1, err := strconv.Atoi(string(r[i1+1:i2]))
	if err != nil {
		return 0, err
	}
	v2, err := strconv.Atoi(string(r[i2+1:i3]))
	if err != nil {
		return 0, err
	}
	return v1 * v2, nil
}