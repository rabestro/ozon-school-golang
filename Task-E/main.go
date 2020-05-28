package main

import (
	"math/rand"
	"sync"
	"time"
)

type Config struct {
	repMutex   *sync.Mutex
	function   func(int) int
	input1     <-chan int
	input2     <-chan int
	output     chan int
	count      int
	repository []int
	wg         *sync.WaitGroup
}

func calc(config *Config, x1 int, x2 int, i int) {
	config.wg.Add(1)
	config.repMutex.Lock()
	config.repository[i] = config.function(x1) + config.function(x2)
	config.repMutex.Unlock()
	config.wg.Done()
}

func process(config *Config) {
	for i := 0; i < config.count; i++ {
		x1, x2 := <-config.input1, <-config.input2
		go calc(config, x1, x2, i)
	}
	config.wg.Wait()

	for i := 0; i < config.count; i++ {
		config.output <- config.repository[i]
	}
	close(config.output)
}

func Merge2Channels(f func(int) int, in1 chan int, in2 chan int, out chan int, n int) {
	task := Config{
		function:   f,
		input1:     in1,
		input2:     in2,
		output:     out,
		count:      n,
		repository: make([]int, n),
		wg:         new(sync.WaitGroup),
		repMutex:   new(sync.Mutex)}

	go process(&task)
}

func fx(x int) int {
	delay := rand.Intn(100000)
	time.Sleep(time.Duration(delay))
	return 2 * x
}

func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 10

	for i := 0; i < n; i++ {

		go func(x int) {
			time.Sleep(time.Duration(rand.Intn(100000)))
			in1 <- x
		} (i)

		go func(x int) {
			time.Sleep(time.Duration(rand.Intn(100000)))
			in2 <- 500 * x
		} (i)

	}

	Merge2Channels(fx, in1, in2, out, n)

	for i := 0; i < n; i++ {
		println(<-out)
	}
}
