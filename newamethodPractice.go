package main

/*

A method is a function with a receiver. 方法是拥有接收器的一种特殊函数
声明格式如下 :
func (r [*]ReceiverType) methodName(param) (result) {
        ...
}

下面 code 为以函数和方法两种办法计算两点之间距离
*/

import (
	"math"
	"fmt"
)

// 声明一个结构体为点 Point 含有横纵左边 X , Y 两个参数 , 参数为 float64 类型
type Point struct {
	X , Y float64
}

// 声明普通函数 Distance
func Distance(p , q Point) float64 {
	return math.Hypot(q.X - p.X , q.Y - p.Y)

}

// 声明 Point 结构体的方法
func (p Point) Distance(q Point ) float64  {
	return math.Hypot(q.X - p.X , q.Y - p.Y)

}

func main()  {
	p := Point{1,2}
	q := Point{4,6}
	fmt.Println(Distance(p, q)) // 调用函数
	fmt.Println(p.Distance(q))	// 调用方法
}

/*

第一个普通函数为包级别的函数 , 称为 geometry.Distance
而第二个方法声明是 Point 结构体的方法 , 所以名字是 Point.Distance
并且由于方法和字段来源于同一个命名空间 , 因此在 Point 结构类型中声明一个叫做 X 的方法会与 X 字段冲突


 */