package main

/*

switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。 
实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 
除非以 fallthrough 语句结束，否则分支会自动终止。 
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
（例如，

switch i {
case 0:
case f():
}

在 i==0 时 f 不会被调用。）

没有条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。

*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var grade string = "B"
	var marks int = getRandInt(5,10)

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 70: grade = "C"
	case 60: grade = "D"
	default: grade = "F"
		
	// 若是想得到区间内的可用if else结构,这里只做switch例子
		
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")

		
	}
	fmt.Printf("你的等级为: %s \n",grade)


}

func getRandInt(min,max int) int {
	rand.Seed(time.Now().Unix()) // 时间戳的种子,很重要,否则只会获得同一个值
	var randmark = ( rand.Intn (max - min) +min ) * 10 // 得到一个 50-90之间的以10为间隔的整值

	fmt.Printf("分数是: %d \n",randmark)
	return  randmark // 返回得到的随机数


