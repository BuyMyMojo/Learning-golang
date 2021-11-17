package main

import "time"

func part1(myChannel chan string) {
	time.Sleep(5 * time.Second)
	myChannel <- "haha"
}

func main() {
	myChannel := make(chan string)

	go part1(myChannel)
	msg := <-myChannel
	println(msg)

}
