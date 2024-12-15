package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var EnableLogging bool

func Log(a ...interface{}) {
	if EnableLogging {
		fmt.Println(a...)
	}
}

func LogSlice(s interface{}) {
	if !EnableLogging { return }
	slice := reflect.ValueOf(s)
	if slice.Kind() != reflect.Slice {
		panic("logSlice() given non-slice argument")
	}
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		Log(item.Interface())
	}
}

func Min(a ...int) int {
	min := math.MaxInt
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(a ...int) int {
	max := math.MinInt
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func AbsDiff(a, b int) int {
	return Abs(a - b)
}

func GetNumbers(s, separator string) []int {
	s = strings.TrimSpace(s)
	res := []int{}
	arr := strings.Split(s, separator)
	for _, v := range arr {
		n, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, n)
		}
	}
	return res
}

func GetIntMat(m, n int) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	return res
}

func GetStrMat(m, n int) [][]string {
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}
	return res
}
