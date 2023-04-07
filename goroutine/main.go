package main

import (
	"fmt"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功:", ret)
}
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启goroutine将0-100的数据发送到ch1中
	go func() {
		for i := 0; i <= 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//开启goroutine从ch1中接受值并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	//在主goroutine中从ch2中接受值打印
	for a := range ch2 {
		fmt.Println(a)
	}

	/*	c := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				c <- i
			}
			close(c)
		}()
		for true {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		fmt.Println("main end...")
	*/
	/*	ch := make(chan int, 1)
		ch <- 10
		fmt.Println("send success")
	*/
	/*	ch := make(chan int)
		go recv(ch)
		ch <- 10
		fmt.Println("发送成功")
	*/
	/*	runtime.GOMAXPROCS(3)
		go a()
		go b()
		time.Sleep(time.Second)
	*/
	/*	go func() {
			i := 0
			for true {
				i++
				fmt.Printf("new goroutine: i = %d\n", i)
				time.Sleep(time.Second)
			}
		}()
		i := 0
		for true {
			i++
			fmt.Printf("main goroutine: i = %d\n", i)
			time.Sleep(time.Second * 2)
			if i == 2 {
				break
			}
		}
	*/
	/*	go func() {
			defer fmt.Println("a.defer")
			func() {
				defer fmt.Println("b.defer")
				runtime.Goexit()
				defer fmt.Println("c.defer")
				fmt.Println("b")
			}()
			fmt.Println("a")
		}()
	*/
}
