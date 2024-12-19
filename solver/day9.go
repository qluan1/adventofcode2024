package solver

import (
	"fmt"
	"strconv"
)

func SolveDay9(input string) error {
	s1, err := d9FirstPart(input)
	if err != nil {
		return err
	}
	fmt.Println("First part:", s1)
	s2, err := d9SecondPart(input)
	if err != nil {
		return nil
	}
	fmt.Println("Second part:", s2)
	return nil
}

func d9FirstPart(input string) (string, error) {
	arr := []int{}
	for i, c := range input {
		v, _ := strconv.Atoi(string(c))
		if i % 2 == 0 { // file
			for j := 0; j < v; j++ {
				arr = append(arr, i/2)
			}
		} else { // free space
			for j := 0; j < v; j++ {
				arr = append(arr, -1)
			}
		}
	}
	l, r := 0, len(arr)-1
	for l < r {
		for arr[l] != -1 && l < r{
			l++
		}
		for arr[r] == -1 && l < r {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	checkSum := 0
	for i, v := range arr {
		if v != -1 {
			checkSum += i * v
		}
	}
	return strconv.Itoa(checkSum), nil
}

func d9SecondPart(input string) (string, error) {
	arr := []int{}
	files := [][]int{}
	spaces := [][]int{}

	for i, c := range input {
		v, _ := strconv.Atoi(string(c))
		if i % 2 == 0 { // file
			tmp := []int{}
			for j := 0; j < v; j++ {
				tmp = append(tmp, i/2)
			}
			files = append(files, tmp)
		} else { // free space
			tmp := []int{}
			for j := 0; j < v; j++ {
				tmp = append(tmp, -1)
			}
			spaces = append(spaces, tmp)
		}
	}
	i := 0
	for i < len(spaces) {
		tmp := spaces[i]
		start := 0
		for start < len(tmp) && tmp[start] != -1 {
			start++
		}
		flag := false
		for j := len(files)-1; j > i; j-- {
			if files[j][0] != -1 && len(tmp) - start >= len(files[j]) {
				flag = true
				for k := 0; k < len(files[j]); k++ {
					tmp[start+k] = files[j][k]
					files[j][k] = -1
				}
				start += len(files[j])
			}
		}
		if !flag || start == len(tmp) {
			i++
		}
	}

	i = 0
	j := 0
	for i < len(files) || j < len(spaces) {
		if i < len(files) {
			arr = append(arr, files[i]...)
			i++
		}
		if j < len(spaces) {
			arr = append(arr, spaces[j]...)
			j++
		}
	}
	checkSum := 0
	for i, v := range arr {
		if v != -1 {
			checkSum += i * v
		}
	}
	return strconv.Itoa(checkSum), nil
}