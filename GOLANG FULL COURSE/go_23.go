package main

// GO ROUTINES
//
//func main() {
//	channel := make(chan string)
//	channel2 := make(chan int)
//
//	go say("Hello, world!", channel, channel2)
//	//time.Sleep(2 * time.Second)
//	fmt.Println(<-channel, <-channel2)
//
//	close(channel)
//	close(channel2)
//}
//
//func say(msg string, ch chan string, ch2 chan int) {
//	fmt.Println(msg)
//
//	ch <- "NO WAY"
//	ch2 <- 2021
//
//}

func main() {
	channel := make(chan int)

	go say("Hello, World!", channel)

	for a := range channel {
		println(a)
	}

	println(<-channel)
}

func say(msg string, ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}
