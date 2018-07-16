package main

import (
	"sort"
	"fmt"
)
// 定义 MyStringList 的类型为 []String
type MyStringList []string
// 实现 sort.Interface 接口的获取元素数量的方法
func (m MyStringList) Len() int  {
	return len(m)
}
// 实现 sort.Interface 接口的比较元素的方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}
// 实现 sort.Interface 接口的交换元素的方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main()  {
	// 随机一个被打乱顺序的字符切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用 sort 包进行排序
	sort.Sort(names)

	// 遍历打印结果
	for _,v := range names{
		fmt.Printf("%s \n", v)

	}
}