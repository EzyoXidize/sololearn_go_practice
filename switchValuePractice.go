package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var grade string = "B"
	var marks int = getRandInt(5,10)

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 70: grade = "C"
	case 60: grade = "D"
	default: grade = "F"
		
	// 若是想得到区间内的可用if else结构,这里只做switch例子
		
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")

		
	}
	fmt.Printf("你的等级为: %s \n",grade)


}

func getRandInt(min,max int) int {
	rand.Seed(time.Now().Unix()) // 时间戳的种子,很重要,否则只会获得同一个值
	var randmark = ( rand.Intn (max - min) +min ) * 10 // 得到一个 50-90之间的以10为间隔的整值

	fmt.Printf("分数是: %d \n",randmark)
	return  randmark // 返回得到的随机数


