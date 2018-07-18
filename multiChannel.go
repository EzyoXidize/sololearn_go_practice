package main

import "fmt"


/*
此方法是通过了建立多个通道进行多个 goroutine 并使 goroutine 的启动按照一定的顺序要求来
此方法适用于过多的 goroutine 时候会导致代码冗长繁琐,容易出错 
*/

func main() {


	ch2 := make(chan bool)
	ch1 := make(chan bool , 1)
	done := make(chan int)

	go func() {
		for i := 1; i <= 50; i++ {
			<-ch1 // 2 , 第 1 步中发送元素到ch1 , ch1 中有元素 , 可接收 否则堵塞
			fmt.Print(2*i - 1,"\t")

			ch2 <- true // 3 , 向 ch2 中发送元素 , 启动另外一个 goroutine
		}
	}()

	go func() {
		for i := 1; i <= 50; i ++ {
			<-ch2 // 4 , 只有当 第一个 goroutine 执行了ch2 中才有元素, 不然一直保持堵塞
			fmt.Println(2 * i)
			ch1 <- true // 5 , 与第一个 goroutine 建立一个通过通道互相启动的机制
		}
		done <- 1	// 6 , 第二个 goroutine 最后向 done 通道发送元素
	}()

	ch1 <- true // 1 , 先往 ch1 中发送一个元素 否则两个 goroutine 均堵塞
	<- done // 7 , 等到第二个 goroutine 循环结束发送元素到 done 通道时候 , 获取元素 ,完成 main goroutine
}



/*

通过3个通道的堵塞机制来建立顺序
利用通道 done 来堵塞主 goroutine 等待其他两个 goroutine 完成
利用通道 ch1 和 ch2 分别互相制约两个协 goroutine
使得两个 goroutine 可以交替进行输出 , 最后完成后主 goroutine 执行完毕
 
Result:
1    2
3    4
...  ...
97    98
99    100

*/
