package main

import "fmt"



/*
 此方法是通过了建立多个通道进行排序goroutine
 */
func main() {


	ch2 := make(chan bool)
	ch1 := make(chan bool , 1)
	done := make(chan int)

	go func() {
		for i := 1; i <= 50; i++ {
			<-ch1 // 2 , ch1 中有元素 , 可接收
			fmt.Print(2*i - 1,"\t")

			ch2 <- true // 3 , 向 ch2 中发送元素
		}
	}()

	go func() {
		for i := 1; i <= 50; i ++ {
			<-ch2 // 4 , 只有当 第一个 goroutine 执行了ch2中才有元素, 不然一直保持堵塞
			fmt.Println(2 * i)
			ch1 <- true // 5 ,
		}
		done <- 1	// 6 , 第二个 goroutine 最后向 done 通道发送元素
	}()

	ch1 <- true // 1 , 先往 ch1 中发送一个元素 否则两个 goroutine 均堵塞
	<- done // 7 , 等到第二个goroutine每次执行完后执行
}