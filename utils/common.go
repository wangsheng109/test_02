package utils

import (
	"log"
	"strconv"
)

func Sum(a, b int) int {

	return a + b

}
func Minus(a, b int) int {
	return a - b
}
func Catch(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	return nums[0] * nums[1] * nums[2]
}

func ToFloat64(str string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[W]", r)
		}
	}()
	if str == "" {
		panic("param is null")
	}

	return strconv.ParseFloat(str, 10)
}
