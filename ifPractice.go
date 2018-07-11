package main

/*

Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。

该语句声明的变量作用域仅在 if 之内。

*/

import (
	"fmt"
	"math"
)

// if 后的条件语句无需加 ()
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // if 中声明的变量 V 同样可以在 else 语句中使用
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//  v 为 if 语句中声明的变量, 不可用在 if 语句之外, 此处不可用 v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}


/*
Result:
27 >= 20 // 此处 Printf 的字符串始终在 return 的值之上
9 20
*/
