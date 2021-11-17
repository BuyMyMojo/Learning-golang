package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	go doStuff1()
	go doStuff2()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}

func doStuff1() {
	i := 1
	for i <= 10 {
		fmt.Println("Thread 1")

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10

		time.Sleep(time.Duration(n) * time.Second)
		i = i + 1
	}
}

func doStuff2() {
	i := 1
	for i <= 10 {
		fmt.Println("Thread 2")

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10

		time.Sleep(time.Duration(n) * time.Second)
		i = i + 1
	}
}
