package main

import (
	"fmt"
	"math/rand"
)
func main() {
	/*
		条件语句：if
		语法格式：
			if 条件表达式{
				//
			}

	 */
	//1.给定一个数字，如果大于10，我们就输出打印这个数字大于10

	// num := 6
	num := rand.Intn(100)
	if num > 10 {
		fmt.Printf("%d  大于10 \n",num)
	}


	//2.给定一个成绩，如果大于等于60分，就打印及格
	// score := 18
	score := num
	
	if score >= 60{
		fmt.Printf("%d成绩及格。\n",score)
	}
	fmt.Println("main..over....")
}
