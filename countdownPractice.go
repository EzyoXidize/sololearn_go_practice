package main
/*
此练习主要是在通过通道和多路复用建立一个根据条件选择继续和中止的程序
利用 abort 通道发送输入的指令, select 语句选择是 <- abort 来判断条件


并利用 goto 语句跳转至结束打印 , 此步骤只为实验 goto 语句 , 对整体代码无关
 */
import (
	"fmt"
	"time"
	"os"
)

func main() {
	fmt.Println("Commencing countdown . Press ENTER to abort.")

	abort := make(chan struct{}) // 创建一个中止通道
	tick := time.Tick(time.Second) // 计时器

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 尝试读取输入的字符
		abort <- struct{}{} // 若有输入则发送至通道 abort 中进行中止语句 , 没输入时候则堵塞
	}()


	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick: // 计时器通道 , 正常工作

		case <-time.After(10 * time.Second):

		case <-abort: // 若通道 abort 中有元素 , 则获取并执行此条内容
			fmt.Println("Launch aborted!")
			goto finish // 跳转至 finish 标签处

		}

	}
	fmt.Println("launch")

	finish:
		fmt.Println("OVER") // 打印中止语句
}



