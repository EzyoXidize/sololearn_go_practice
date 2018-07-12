package main

import "fmt"

type students struct {
	// 首先声明一个结构体 students , 其中包含有 name ,sex grade ,class ,id 五个属性以及各自的类型
	name string
	sex string
	grade int
	class int
	id int
}

func main()  {
	var student1 students	// 声明 student1 为 students 类型
	var student2 students	// 声明 student2 为 students 类型

	student1.name = "lilei"
	student1.sex = "boy"
	student1.grade = 3
	student1.class = 2
	student1.id = 112011245

	student2.name = "hanmeimei"
	student2.sex = "girl"
	student2.grade = 3
	student2.class = 1
	student2.id = 112111245

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