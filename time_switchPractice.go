package main

import (
	"fmt"
	"time"  // import time 库
)

func main(){
	i := 2
	fmt.Print("write ",i," as ") // switch case 语句 practice
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:  // 用 switch  case 语句判断条件
		fmt.Println("its the weekend")
	default:			// 默认值
		fmt.Println("its the weekday")

	}

	t := time.Now()
	switch   {
	case t.Hour() < 12:
		fmt.Println("its the morning")
	default:
		fmt.Println("its the afternoon")
		
	
	}
}

/*

返回的结果举例

write 2 as two
its the weekday
its the morning

*/
