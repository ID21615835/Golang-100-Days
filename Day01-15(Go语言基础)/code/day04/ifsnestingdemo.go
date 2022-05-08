package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	/*
	if语句的嵌套：
	if 条件1{
		A段
	}else{
		if 条件2{
			B段
		}else{
			C段
		}
	}

	简写：
	if 条件1{
		A段 //条件1成立
	}else if 条件2{
		B段 // 条件1不成立，条件2成立
	}else if 条件3{
		C段 // 条件1不成立，条件2不成立，条件3成立。。

	}.。。else{
	}

	 */
	
	t1 := time.Now()

	for i :=0; i <= 19900; i++ {

		wc();
	}  
	
	t2 := time.Now()

	fmt.Println("t1与t2 时间差:",t2.Sub(t1))
	fmt.Println("main...over....")
}

func wc() {

	listsex := [...] string{"男","女","其他"}
	// sex := "泰国" //bool, int, string
	hex := rand.Intn(3)
	fmt.Printf("hex 为：%d \n" , hex)
	sex := listsex[hex]

	if sex == "男" {
		fmt.Println("可以去男厕所啦。。。") // sex 是男
	} else if sex == "女" {
		fmt.Println("你去女厕所吧。。") //sex 是女

	} else {
		fmt.Println("我也不知道了。。")

	}

}