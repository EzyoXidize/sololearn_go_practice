package main

import "fmt"

/*

go 中结构体是 由一系列具有相同类型或不同类型的数据构成的数据集合 。
结构体定义需要 type 和 struct 语句 , 如
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}

结构体成功声明后 , 就可用于变量的声明
访问结构体某一个内部构件时候 , 用 "." 符号 ,操作方式为 结构体.构件 如 student1.name 

*/

type students struct {     // 首先声明一个结构体 students , 其中包含有 name ,sex grade ,class ,id 五个构件以及各自的类型
	name string
	sex string
	grade int
	class int
	id int
}

func main()  {
	var student1 students   // 声明 student1 为 students 类型
	var student2 students   // 声明 student2 为 students 类型
        
        // student1 描述
	student1.name = "lilei"
	student1.sex = "boy"
	student1.grade = 3
	student1.class = 2
	student1.id = 112011245
        
        // student2 描述
	student2.name = "hanmeimei"
	student2.sex = "girl"
	student2.grade = 3
	student2.class = 1
	student2.id = 112111245
        
        // 打印 student1 和 student2 的信息
	printStudent(student1)
	fmt.Printf("\n")
	printStudent(student2)


}

func printStudent(student students)  {
	fmt.Printf("student's name : %s \n",student.name)
	fmt.Printf("student's sex : %s \n",student.sex)
	fmt.Printf("student's grade : %d \n",student.grade)
	fmt.Printf("student's class : %d \n",student.class)
	fmt.Printf("student's id : %d \n",student.id)
}

/*

Result:

student's name : lilei 
student's sex : boy 
student's grade : 3 
student's class : 2 
student's id : 112011245 

student's name : hanmeimei 
student's sex : girl 
student's grade : 3 
student's class : 1 
student's id : 112111245 

*/
