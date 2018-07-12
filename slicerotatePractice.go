package main

import "fmt"

/*

习题: 声明一个函数用来 rotate , 实现一次遍历完成元素旋转
思路: 利用一个额外的盒子 slice , 从元 slice 里倒序获取元素后存入新 slice
疑问: 效率问题,以及是否有更好的办法

 */

func main()  {
	x := []int{1,2,3,4,5}
	x = rotate(x)
	fmt.Println(x)

}

var y []int // 声明一个用来 append 转换元素的新 slice y

func rotate(x []int ) []int {	// 声明 ratate 函数倒序调取 slice x 里的元素并存入 y
	for i := 0; i < 5; i++ {

		y = append(y, x[4-i])

	}
	return y
}