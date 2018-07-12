package main

import "fmt"

func main()  {
	var a int = 10
	var ip *int // 声明指针变量为 整型指针
        var ptr *int

	ip = &a 	// 赋值指针变量 ip 变量 a 的地址

	// 输出 a 变量的地址
	fmt.Printf("a 变量的地址为 : %d \n",&a)
	// 输出指针变量指向的存储地址
	fmt.Printf("a 变量的地址为 : %d \n",ip)
	// 输出指针指向的值
	fmt.Printf("*ip 变量的值为 : %d \n",*ip)
        // 输出空指针的值
        fmt.Printf("输出空指针 ptr 的值 : %d \n",ptr)


}


/*
Resul:

a 变量的地址为 : 825741050040 
a 变量的地址为 : 825741050040 
*ip 变量的值为 : 10 
输出空指针 ptr 的值 : 0 

空指针判断:
if(ptr != nil)    不是空指针 
if(ptr == nil)    是空指针

*/
