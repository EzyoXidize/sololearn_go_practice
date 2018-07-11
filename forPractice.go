package main

/*

Go语言中只有for循环一种循环语句
for循环语句的语法和其他语言的类似(但是Go中在for语句后不需要加 小括号() 而大括号{}是必须的),基本的for循环由三个部分组成,并用 ":" 隔开
    for 初始化语句 : 条件表达式 : 后置语句
    
    初始化语句：在第一次迭代前执行
    条件表达式：在每次迭代前求值
    后置语句：在每次迭代的结尾执行
    
一旦条件表达式为 false ,则终止循环

*/

import "fmt"

func main() {
	// 声明变量 a 和 b (赋值为 5 )
	var a int
	var b int = 5

	// 定义一个 arr , size 为 6 
	nums := [3]int{0,1,3}

	// for 循环语句
	for a := 0 ; a < 5; a ++ {
		fmt.Println("a is",a)

	}

	for a < b {
		a ++
		fmt.Println("a is",a)

	}

	for i,x := range nums {
		fmt.Printf("第 %d 位的值为 %d \n",i,x)

	}
}





/*
Result:

a is 0
a is 1
a is 2
a is 3
a is 4
a is 1
a is 2
a is 3
a is 4
a is 5
第 0 位的值为 0
第 1 位的值为 1
第 2 位的值为 3

*/
