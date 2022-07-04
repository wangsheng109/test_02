package main

import (
	"fmt"

	"test_02/utils"
)

func main() {

	fmt.Printf("I Love You forever\n")
	c := utils.Sum(1, 2)
	fmt.Printf("a+b=%d\n", c)
	d := utils.Minus(1, 2)
	fmt.Printf("a-b=%d\n", d)
	utils.Test(1, 2)
	//err := "程序异常退出"
	//log.Fatal(err)
	//log.Panicf("%s", err)
	utils.Catch(2, 8)
	utils.ToFloat64("123")
}
