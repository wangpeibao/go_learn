package main

import "fmt"

// goroutine + channel
// channel是goroutine间的通信方式

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
	fmt.Println("Counting")
}

func main() { // 并发编程
	chs := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	// channel的定义方式 var chanName chan ElementType
	var ch1 chan int
	ch1 <- 1
	fmt.Println(<-ch1)
}
