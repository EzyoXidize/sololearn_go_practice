package main

import "fmt"

/*

Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
for [i/_], iterator := range [array_name/slice_name/channel_name/map_name]
i/_ : 需要使用时可以用 i 当下标，不需要时直接用 _ 代替
array_name/slice_name/channel_name/map_name : 多选一

*/


func main()  {
	// 利用range求一个slice的和, 若是用数组方法可类似,此处不需要用到索引key的值,所以前面用 "_" 省略掉了,需要索引值时候下面例子
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums { // _,num _为空白占位 , 忽略
		sum += num
	}

	fmt.Println("sum = ",sum)

	// 需要知道索引值时候, 用range 传入 index和值两个变量
	for i,num := range nums{
		if num == 3 {
			fmt.Printf("index: %d \n",i)
			
		}
	}
	 // range也可用在 map的 key - value对上
	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s --> %s \n",k,v)

	}

	// range也可用来枚举 Unicode字符串,第一个参数为字符索引,第二个参数是字符(Unicode的值) 本身
	for i,c := range "go" {
		fmt.Println(i,c)
	}

}


/*

Result:

sum =  9
index: 1 
a --> apple 
b --> banana 
0 103
1 111

*/
