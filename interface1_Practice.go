package main

import "fmt"

/*
声明接口格式:
type Interface_Name interface {
    method_a() string
    method_b() int ....
}
方法名首字母是大写时候 , 该方法可被接口所在包之外的代码访问
 */

// 声明一个数据写入器
type dataWriter interface {
	writeData(data interface{}) error
}

// 声明文件结构
type file struct {
}

// 实现 DateWriter 接口的 WriteData() 方法
func (d *file) writeData(data interface{}) error {

	// 模拟写入数据 , 相当于为 file 定义了一个方法 , 绑定了接口里的一个方法 writeData()
	fmt.Println("WriteData:", data)
	return nil

}

func main()  {

	// 实例化 file
	f := new(file)

	// 声明一个 DataWriter 的接口
	var writer dataWriter

	// 将接口赋值 f , 也就是 *file 类型
	writer = f

	// 使用 DateWriter 接口进行数据写入
	writer.writeData("try to write something")

}

